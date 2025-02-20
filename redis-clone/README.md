# Redis Clone in Go

A minimal Redis server implementation in Go that handler client communication using Redis Serialization Protocol ([RESP](https://redis.io/topics/protocol)).

## Project Structure

```
redis-clone/
├── cmd/
│   └── server/
│       └── main.go    # TCP server implementation
├── resp/
│   └── resp.go        # RESP protocol serialisation/de-serialisation implementation
└── README.md
```

## Features

- TCP server implementation on Redis default port (6379)
- Handles client connections concurrently
- Parse and print received client commands on the console
- Currently responds with "OK" to all requests

## Getting Started

1. Clone the repository:
```bash
git clone https://github.com/faizan2786/go-projects.git
cd projects/redis-clone
```

2. Run the server:
```bash
go run cmd/server/.
```

3. Test using redis-cli:
```bash
redis-cli
> ping
OK
> set key value
OK
```

## Implementation Details

- Server listens on port 6379 (Redis default port)
- Uses Go's net package for TCP communication
- Implements basic RESP protocol deserialisation
- Handles concurrent client connections using goroutines

## Future Enhancements

- [ ] RESP writer for server responses
- [ ] Support more RESP data types
- [ ] Implement key-value storage
- [ ] Add data persistence

## License

MIT
