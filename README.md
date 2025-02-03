# Go Projects Collection

This repository contains various Go projects demonstrating different aspects of Go programming.

## Projects

### RPC
A demonstration of Remote Procedure Calls (RPC) in Go using the standard `net/rpc` package. Features include:
- Basic RPC server and client implementations
- Client with Timeout handling
- Graceful shutdown for server
- RPC examples with multiple services

[View RPC Project →](./RPC)

### Project Structure
```
projects/
├── RPC/                # RPC demonstration project
│   ├── services/       # RPC service definitions
│   ├── server/         # Server implementation
│   ├── main/           # Client implementations
│   └── README.md       # Project documentation
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
