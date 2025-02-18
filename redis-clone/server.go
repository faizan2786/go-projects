package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
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

	// deserialise and print the received data from the client
	fmt.Println("Client sent:", deserialiseClientCommand(data))

	// ignore the request: always respond clients with a fixed string: "OK"
	resp := []byte("+OK\r\n") // return a plain string in Redis Serialization Protocol
	_, err = conn.Write(resp)
	if err != nil {
		fmt.Println("Error while writing to client:", err)
		return
	}
}

// define a list of characters for RESP data type code
type RESPCode byte

const (
	BulkString RESPCode = '$' // character code for bulk string type data
	Array      RESPCode = '*'
)

// convert bytes (data) received from the client to a string
// (For simplicity, we only expect to receive a bulk string type with a single digit size data only)
func deserialiseClientCommand(data []byte) string {

	// create a buffer reader to read string data
	reader := bufio.NewReader(strings.NewReader(string(data)))

	// skip the first array type code
	// (Every command sent by redis-cli is of type array. We are interested in the data which resided in the first element of this array)
	reader.ReadBytes('\n')

	// now, parse the actual command...

	// read the next byte and verify if it a bulk string type
	b, _ := reader.ReadByte()

	if b != byte(BulkString) {
		fmt.Println("Invalid first byte of the data, expecting first the code for bulk strings only.")
		fmt.Println("Actual data received:", string(data))
		os.Exit(1)
	}

	// read the size of the string (limiting to 1 digit size only )
	b, _ = reader.ReadByte()
	strSize := int(b - '0') // convert ASCII byte code to its int value

	// skip next two bytes for \r\n
	reader.ReadByte()
	reader.ReadByte()

	// read the next bytes of size = strSize
	bytes := make([]byte, strSize)
	reader.Read(bytes)
	return string(bytes) // convert the read bytes into string
}
