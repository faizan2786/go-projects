package resp

import (
	"bufio"
	"fmt"
	"io"
)

// This file defines RESP reader and writer
// A RESP reader is used to deserialize the messages received from clients
// in [RESP](https://redis.io/docs/latest/develop/reference/protocol-spec/#resp-protocol-description) format

// RESP type codes
type RespCode byte

const (
	ARRAY RespCode = '*'
	BULK  RespCode = '$' // bulk string type
)

type RespIO struct {
	reader *bufio.Reader
}

// accepts any object that implements io.Reader interface (i.e. net package's Conn object)
// returns new RespIO object
func NewRespIO(rd io.Reader) *RespIO {
	return &RespIO{reader: bufio.NewReader(rd)}
}

// deserialize bytes (data) from the io Reader
// (For simplicity, we only expect only a bulk string type with a single digit size data only)
func (r *RespIO) ReadCommand() (string, error) {

	// Every command sent by redis-cli is of type Array.
	// We are interested in the data which resided in the first element of this array

	// Read the first byte
	_, err := r.reader.ReadByte()
	if err != nil {
		return "", err
	}

	// skip till the first array element
	r.reader.ReadBytes('\n')

	// now, parse the actual command...

	// read the next byte and verify if it a bulk string type
	b, _ := r.reader.ReadByte()

	if b != byte(BULK) {
		return "", fmt.Errorf("invalid first byte, expecting byte for bulk strings only.Actual byte received:%s", string(b))
	}

	// read the size of the string (limiting to 1 digit size only )
	b, _ = r.reader.ReadByte()
	strSize := int(b - '0') // convert ASCII byte code to its int value

	// skip next two bytes for \r\n
	r.reader.ReadByte()
	r.reader.ReadByte()

	// read the next bytes of size = strSize
	bytes := make([]byte, strSize)
	r.reader.Read(bytes)

	return string(bytes), nil
}
