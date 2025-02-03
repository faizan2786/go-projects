# Go RPC Demo Project

This project demonstrates the implementation of Remote Procedure Calls (RPC) in Go using the standard `net/rpc` package.

## Project Structure

```
RPC/
├── services/
│   └── demo.go         # RPC service definitions
├── server/
│   └── rpcserver.go    # RPC server implementation
├── main/
│   ├── client.go                # Basic client implementation
│   └── client_with_timeout.go   # Client with timeout handling
└── README.md
```

## Features

- Basic RPC server implementation
- Two client implementations:
  - Basic client with synchronous calls
  - Advanced client with timeout handling
- Graceful server shutdown
- Demo RPC service with example methods
- TCP communication over HTTP

## Getting Started

1. Clone the repository:
```bash
git clone https://github.com/faizan2786/go-projects.git
cd RPC
```

2. Start the server:
```bash
go run main/server.go
```

3. Run either client (in a separate terminal):

Basic client:
```bash
go run main/client.go
```

Client with timeout:
```bash
go run main/client_with_timeout.go
```

## Implementation Details

- Server runs on localhost:1298
- Uses context for timeout handling
- Implements graceful shutdown with OS signal handling
- Demonstrates both simple and long-running RPC calls

## Available RPC Methods

- `RPCDemoService.GetServerMessage`: Returns a simple message from the server
- `RPCDemoService.SomeLongRunningProcess`: Simulates a long-running process (5 seconds)
- `Arith.Multiply`: Multiplies two numbers
- `Arith.Divide`: Divides two numbers, returns quotient and remainder

## License

MIT
