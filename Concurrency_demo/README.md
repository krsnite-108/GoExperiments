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

### Concurrent Implementation

File: `Smart_PJ_API.go`

This version uses goroutines to handle multiple API requests simultaneously, reducing the execution time to 1.36 seconds.