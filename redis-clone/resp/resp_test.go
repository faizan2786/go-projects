package resp

import (
	"strings"
	"testing"
)

func TestReadMessage(t *testing.T) {

	// a test struct
	// (containing a name, input string, expected output, expected error flag and error message)
	type testCase struct {
		name    string
		input   string
		want    []string
		wantErr bool
		errMsg  string
	}

	// a slice of test structs - each element represents a test-case
	tests := []testCase{
		{
			name:    "simple PING command",
			input:   "*1\r\n$4\r\nPING\r\n",
			want:    []string{"PING"},
			wantErr: false,
		},
		{
			name:    "SET command with key and value",
			input:   "*3\r\n$3\r\nSET\r\n$3\r\nkey\r\n$5\r\nvalue\r\n",
			want:    []string{"SET", "key", "value"},
			wantErr: false,
		},
		{
			name:    "mixed types array",
			input:   "*3\r\n$3\r\nGET\r\n:123\r\n+Hello\r\n",
			want:    []string{"GET", "123", "Hello"},
			wantErr: false,
		},
		{
			name:    "array with empty bulk string",
			input:   "*1\r\n$-1\r\n",
			want:    []string{""},
			wantErr: false,
		},
		{
			name:    "invalid first byte",
			input:   "$1\r\n1\r\n",
			wantErr: true,
			errMsg:  "invalid first byte",
		},
		{
			name:    "invalid array length",
			input:   "*abc\r\n",
			wantErr: true,
			errMsg:  "invalid syntax",
		},
		{
			name:    "incomplete bulk string",
			input:   "*1\r\n$5\r\nHEL",
			wantErr: true,
			errMsg:  "EOF",
		},
		{
			name:    "unsupported type",
			input:   "*1\r\n!5\r\nHELLO\r\n",
			wantErr: true,
			errMsg:  "unsupported RESP type code",
		},
	}

	// define a test function
	// this function will be called by t.Run for each test case
	testFunc := func(t *testing.T, tt testCase) {

		// strings.NewReader(tt.input) creates a new Reader that reads from a string
		// It converts the test input string into something that behaves like a stream of data
		// This simulates how the actual Redis protocol would receive data over a network connection
		// (i.e. instead of reading from a network socket, we're reading from a string in memory)

		reader := NewRespIO(strings.NewReader(tt.input))
		got, err := reader.ReadMessage()

		// Check errors first (i.e. for invalid cases)
		if tt.wantErr {

			// no error occurred, but we were expecting one
			if err == nil {
				t.Errorf("ReadMessage() error = nil, wantErr %v", tt.wantErr)
				return
			}

			// error occurred, but the error message doesn't match what we were expecting
			if tt.errMsg != "" && !strings.Contains(err.Error(), tt.errMsg) {
				t.Errorf("ReadMessage() error = %v, want %v", err, tt.errMsg)
			}

			return
		}

		// Check success cases
		if err != nil {
			t.Errorf("ReadMessage() unexpected error = %v", err)
			return
		}

		if len(got) != len(tt.want) {
			t.Errorf("ReadMessage() got %v elements, want %v elements", len(got), len(tt.want))
			return
		}

		for i := range got {
			if got[i] != tt.want[i] {
				t.Errorf("ReadMessage() got[%d] = %v, want[%d] = %v", i, got[i], i, tt.want[i])
			}
		}
	}

	// run the test cases (i.e. for each test struct, run the test function)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) { testFunc(t, tt) })
	}
}

func TestReadBulkString(t *testing.T) {

	type testCase struct {
		name    string
		input   string
		length  int
		want    string
		wantErr bool
	}

	tests := []testCase{
		{
			name:    "valid bulk string",
			input:   "Hello\r\n",
			length:  5,
			want:    "Hello",
			wantErr: false,
		},
		{
			name:    "empty bulk string",
			input:   "\r\n",
			length:  0,
			want:    "",
			wantErr: false,
		},
		{
			name:    "incomplete bulk string",
			input:   "Hel",
			length:  5,
			wantErr: true,
		},
	}

	testFunc := func(t *testing.T, tt testCase) {

		reader := NewRespIO(strings.NewReader(tt.input))
		got, err := reader.readBulkString(tt.length)

		if tt.wantErr {
			if err == nil {
				t.Errorf("readBulkString() error = nil, wantErr %v", tt.wantErr)
			}
			return
		}

		if err != nil {
			t.Errorf("readBulkString() unexpected error = %v", err)
			return
		}

		if got != tt.want {
			t.Errorf("readBulkString() = %v, want %v", got, tt.want)
		}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) { testFunc(t, tt) })
	}
}
