package main

import (
	"encoding/json" // For encoding/decoding into JSON
	"fmt"
	"io"       // For reading from responses
	"net/http" // For making HTTP requests
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

func main() {
	// Create a custom HTTP client
	client := &http.Client{}

	// Create a new GET request to the Dad Jokes API
	request, err := http.NewRequest("GET", "https://icanhazdadjoke.com/", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Add headers to the request
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")
	fmt.Println("Fetching jokes from the Dad Jokes API...")

	// Track the start time
	startTime := time.Now()

	// Fetch data from the API 100 times
	for i := 0; i < 100; i++ {
		response, err := client.Do(request)
		if err != nil {
			fmt.Println("Error making request:", err)
			continue
		}
		defer response.Body.Close() // Ensure the response body is closed

		// Read the response body
		msgBody, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			continue
		}

		// Unmarshal the JSON response into the APIResponse struct
		var responseObject APIResponse
		err = json.Unmarshal(msgBody, &responseObject)
		if err != nil {
			fmt.Println("Error unmarshaling JSON:", err)
			continue
		}

		// Print the joke message
		fmt.Println(responseObject.Message)
	}

	// Calculate and print the total elapsed time
	elapsedTime := time.Since(startTime)
	fmt.Printf("Total time taken: %s\n", elapsedTime)
}
