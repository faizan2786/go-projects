package resp

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

// This file defines RESP reader and writer
// A RESP reader is used to deserialize the messages received from clients
// in [RESP](https://redis.io/docs/latest/develop/reference/protocol-spec/#resp-protocol-description) format

// RESP type codes
type RespCode byte

const (
	ARRAY   RespCode = '*'
	BULK    RespCode = '$' // bulk string type
	INTEGER RespCode = ':'
	STRING  RespCode = '+' // simple string type
)

type RespIO struct {
	reader *bufio.Reader
}

// accepts any object that implements io.Reader interface (i.e. net package's Conn object)
// returns new RespIO object
func NewRespIO(rd io.Reader) *RespIO {
	return &RespIO{reader: bufio.NewReader(rd)}
}

// deserialize RESP message into a set of strings
func (r *RespIO) ReadMessage() ([]string, error) {

	// Every command sent by redis-cli is of type Array.
	// We are interested in the data which resided in the first element of this array

	// Read the first byte
	b, err := r.reader.ReadByte()
	if err != nil {
		return nil, err
	}

	if b != byte(ARRAY) {
		return nil, fmt.Errorf("invalid first byte, expecting byte for array type ('*') only. Actual byte received:%s", string(b))
	}

	return r.readArray()
}

// read a RESP array
func (r *RespIO) readArray() ([]string, error) {

	values := make([]string, 0)

	numElements, err := r.readInt()
	if err != nil {
		return nil, err
	}

	// read each element of the array
	for i := 0; i < numElements; i++ {

		// check if the next element is a bulk string
		b, _ := r.reader.ReadByte()

		var val string
		var err error
		switch RespCode(b) {
		case BULK:
			// read the length of the string
			var length int
			length, err = r.readInt()
			if err != nil {
				return nil, err
			}
			// read the string
			val, err = r.readBulkString(length)
		case INTEGER:
			// read the integer
			var intVal int
			intVal, err = r.readInt()
			if err == nil {
				val = strconv.Itoa(intVal)
			}
		case STRING:
			// read the string
			val, err = r.readLine()
		default:
			return nil, fmt.Errorf("unsupported RESP type code: %v", b)
		}

		if err != nil {
			return nil, err
		}

		values = append(values, val)
	}

	return values, nil
}

// read the next integer
func (r *RespIO) readInt() (int, error) {
	strNum, err := r.readLine()
	if err != nil {
		return 0, err
	}

	num, err := strconv.ParseInt(strNum, 10, 64)
	if err != nil {
		return 0, err
	}
	return int(num), nil
}

// read the RESP buffer up to the next line (i.e. next \r)
// returns the line as string (without \r\n)
func (r *RespIO) readLine() (string, error) {
	line, err := r.reader.ReadString('\r')
	if err != nil {
		return "", err
	}

	// skip the next byte (i.e. \n)
	r.reader.ReadByte()

	// strip the trailing \r from the line
	return line[:len(line)-1], nil
}

// read a bulk string of the specified length
func (r *RespIO) readBulkString(length int) (string, error) {
	// create a buffer to read the bulk string
	bulk := make([]byte, length)

	// read the string
	_, err := r.reader.Read(bulk)
	if err != nil {
		return "", err
	}

	// skip CRLF after the bulk string
	r.reader.ReadByte()
	r.reader.ReadByte()

	return string(bulk), nil
}
