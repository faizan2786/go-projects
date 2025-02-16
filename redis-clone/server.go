package main

import (
	"fmt"
	"net"
)

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

	// read the data from client
	data := make([]byte, 1024)
	_, err := conn.Read(data)
	if err != nil {
		fmt.Println("Error while receiving data from the client")
		return
	}
	// fmt.Println(string(data)) // print the received data

	// ignore the request: always respond clients with a fixed string: "OK"
	resp := []byte("+OK\r\n") // return a plain string in Redis Serialization Protocol
	_, err = conn.Write(resp)
	if err != nil {
		fmt.Println("Error while writing to client:", err)
		return
	}
}
