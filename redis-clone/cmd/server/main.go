package main

import (
	"fmt"
	"net"

	"github.com/faizan2786/go-projects/redis-clone/resp"
)

// A basic TCP server that listens on Redis default port (6379)
// Accepts client connections,
// read their commands in RESP protocol format
// and responds with "OK" (in RESP) to all requests

const port string = "6379" // redis server default port

func main() {

	// start the default TCP listener
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Error while starting the server:", err)
		return
	}

	// keep listening for client connections
	fmt.Printf("Server is ready to accept connections on port %s...\n", port)
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error while accepting client connection:", err)
			return
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	reader := resp.NewRespIO(conn)

	// deserialize and print the received data from the client
	commands, err := reader.ReadMessage()

	if err != nil {
		fmt.Println("Error while reading from client:", err)
		return
	}

	fmt.Println("Client sent:", commands)

	// ignore the request: always respond clients with a fixed string: "OK"
	resp := []byte("+OK\r\n") // return a plain string in Redis Serialization Protocol
	_, err = conn.Write(resp)
	if err != nil {
		fmt.Println("Error while writing to client:", err)
		return
	}
}

// Note: Instead of "go run main.go", run "go run ." to compile all package files.
