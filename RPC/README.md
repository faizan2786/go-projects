# Go RPC Demo Project

This project demonstrates the implementation of Remote Procedure Calls (RPC) in Go using the standard `net/rpc` package.

## Project Structure

```
RPC/
├── cmd/
│   ├── client/
│   │   └── main.go                # Basic client implementation
│   ├── client_with_timeout/
│   │   └── main.go                # Client with timeout handling
│   └── server/
│       └── main.go                # RPC server implementation
├── services/
│   └── demo.go                    # RPC service definitions
│   └── arith.go                   # Arithmetic service definitions
│   └── types.go                   # Argument type definition
└── README.md
```

## Features

- Basic RPC server implementation with graceful shutdown
- Two services with example RPC methods
- Two client implementations:
  - Basic client with synchronous calls
  - Advance client with asynchronous call and timeout handling

## Getting Started

Clone the repository:
```bash
git clone https://github.com/faizan2786/go-projects.git
cd projects/RPC
```

Start the server:
```bash
go run cmd/server
```

Run either client (in a separate terminal):

- For basic client:
  ```bash
  go run cmd/client
  ```

- For client with timeout:
  ```bash
  go run cmd/client_with_timeout
  ```

You may also **first compile** the server and client binaries:
```bash
go build cmd/server/
go build cmd/client/
go build cmd/client_with_timeout/
```

Then **run** the binaries:
```bash
./server
./client
./client_with_timeout
```

## Implementation Details

- Server runs on localhost:1298
- Uses context for timeout handling
- Implements graceful shutdown with OS signal handling
- Clients demonstrate both simple and long-running RPC calls

## Available RPC Methods

- `RPCDemoService.GetServerMessage`: Returns a simple message from the server
- `RPCDemoService.SomeLongRunningProcess`: Simulates a long-running process (5 seconds)
- `Arith.Multiply`: Multiplies two numbers
- `Arith.Divide`: Divides two numbers, returns quotient and remainder

## License

MIT
