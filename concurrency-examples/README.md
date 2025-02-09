
# Go Concurrency Examples

This project folder contains practical examples demonstrating application of various concurrency patterns in Go.

## Examples

### Event Loop (`event_loop.go`)
Implements an asynchronous event processing system:
- Continuous event loop listening for requests
- Async event handling with callbacks
- Graceful shutdown mechanism
- Uses buffered channels for event queueing

### Load Balancer (`load_balancer.go`)
Demonstrates a worker pool pattern where multiple workers process requests concurrently:
- Uses channels for request distribution
- Implements fan-out/fan-in pattern
- Shows dynamic work distribution among workers

### Replicated Servers (`replicated_servers.go`)
Shows how to handle redundant server requests:
- Sends parallel requests to multiple servers
- Returns fastest response using select
- Implements timeout pattern
- Simulates real-world distributed systems scenario

## Running the Examples

Each example can be run independently using:

```bash
go run <filename>.go
```

For instance:
```bash
go run load_balancer.go
```

## Key Concepts Demonstrated

- Goroutines & Channels (buffered and unbuffered)
- Select statement
- Timeout pattern
- Distributed Workers
- Fan-out/Fan-in pattern