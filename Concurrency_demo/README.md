# Performance Improvement with Goroutines

This repository demonstrates the performance gains achieved by using Go's goroutines for concurrent execution.

## Performance Comparison

- **Single-threaded Execution (Poor_PJ_API.go)**: 38 seconds
- **Concurrent Execution with Goroutines (Smart_PJ_API.go)**: 1.36 seconds

## How It Works

- **Single-threaded Execution** (`Poor_PJ_API.go`): Requests are processed sequentially, resulting in a longer total execution time.
- **Concurrent Execution** (`Smart_PJ_API.go`): Multiple requests are processed in parallel using goroutines, significantly reducing the total execution time.

## Code Overview

### Single-threaded Implementation

File: `Poor_PJ_API.go`

This version processes API requests one at a time, leading to a total execution time of 38 seconds.
![Screenshot 2024-09-06 at 11 53 50 PM](https://github.com/user-attachments/assets/ba1f7c05-701c-4f9c-bcf6-35f2c5fddf5f)

### Concurrent Implementation

File: `Smart_PJ_API.go`

This version uses goroutines to handle multiple API requests simultaneously, reducing the execution time to 1.36 seconds.
![Screenshot 2024-09-07 at 12 11 36 AM](https://github.com/user-attachments/assets/15b13757-a371-41e4-8628-bef976cfff16)
