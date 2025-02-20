# Go Projects Collection

This repository contains various Go projects demonstrating different aspects of Go programming.

## Projects

### Redis Clone
A basic Redis server clone that:
- Implements TCP server on Redis default port (6379)
- Accepts multiple client connections
- Handles RESP (Redis Serialization Protocol) for client communication

[View Redis Clone Project →](./redis-clone)

### RPC
A demonstration of Remote Procedure Calls (RPC) in Go using the standard `net/rpc` package. Features include:
- Basic RPC server and client implementations
- Client with Timeout handling
- Graceful shutdown for server using `Cancellable Context`
- RPC examples with multiple services

[View RPC Project →](./RPC)

### Concurrency Examples
Collection of practical concurrency patterns in Go, including:
- Event Loop implementation
- Load Balancer with multiple workers
- Replicated Servers with timeout handling
- Demonstrates various Go concurrency primitives

[View Concurrency Examples →](./concurrency-examples)

### Project Structure
```
projects/
│
├── redis-clone/        # Redis server clone project
│   ├── cmd/            # Executables
│   │   └── server/     # TCP Server implementation
│   ├── resp/           # RESP protocol serialisation/de-serialisation
│   └── README.md       # Project documentation
│
├── RPC/                            # RPC demonstration project
│   ├── cmd/                        # Executables
│   │   ├── client/                 # Basic client implementation
│   │   ├── client_with_timeout/    # Client with timeout handling
│   │   └── server/                 # RPC server implementation
│   ├── services/                   # RPC service definitions
│
├── concurrency-examples/
│   ├── event_loop.go
│   ├── load_balancer.go
│   ├── replicated_servers.go
│   └── README.md
│
└── README.md           # This file
```

## Getting Started

Each project has its own README with specific setup and running instructions. To get started:

1. Clone the repository:
```bash
git clone https://github.com/faizan2786/go-projects.git
```

2. Navigate to the desired project directory:
```bash
cd projects/<project-name>
```

3. Follow the project-specific README instructions.

## Requirements

- Go 1.22 or higher

## License

MIT
