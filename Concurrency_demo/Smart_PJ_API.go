package main

import (
	"encoding/json" // For encoding/decoding into JSON
	"fmt"
	"io"       // For reading from responses
	"net/http" // For making HTTP requests
	"sync"     // For synchronizing goroutines
	"time"
)

// APIResponse models the API response from https://icanhazdadjoke.com
// The JSON tags specify how struct fields should be named in the resulting JSON
// when the struct is converted to JSON format (marshaled).
type APIResponse struct {
	ID           string `json:"id"`
	Message      string `json:"joke"`
	ResponseCode int    `json:"status"`
}

// Declare a global WaitGroup to manage goroutine synchronization
var wg sync.WaitGroup

func main() {
	// Set the number of goroutines to wait for
	wg.Add(100)

	fmt.Println("Fetching jokes from the Dad Jokes API...")

	// Record the start time
	startTime := time.Now()

	// Launch 100 goroutines to fetch jokes
	for i := 0; i < 100; i++ {
		go getJoke()
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// Calculate and print the total elapsed time
	elapsedTime := time.Since(startTime)
	fmt.Printf("\n\n\nTotal time taken: %s\n", elapsedTime)
}

// getJoke makes an HTTP request to fetch a joke and prints it.
// It uses the global WaitGroup to signal when the goroutine is done.
func getJoke() {
	client := &http.Client{}

	// Create a new GET request
	request, err := http.NewRequest("GET", "https://icanhazdadjoke.com/", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		wg.Done() // Signal that this goroutine is done, even if there was an error
		return
	}

	// Set request headers
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")

	// Send the request and get the response
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error making request:", err)
		wg.Done() // Signal that this goroutine is done, even if there was an error
		return
	}
	defer response.Body.Close() // Ensure the response body is closed

	// Read the response body
	msgBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		wg.Done() // Signal that this goroutine is done, even if there was an error
		return
	}

	// Unmarshal the JSON response into the APIResponse struct
	var responseObject APIResponse
	err = json.Unmarshal(msgBody, &responseObject)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		wg.Done() // Signal that this goroutine is done, even if there was an error
		return
	}

	// Print the joke message
	fmt.Println(responseObject.Message)

	// Signal that this goroutine is done
	wg.Done()
}
