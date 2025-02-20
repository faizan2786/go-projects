package services

import (
	"math/rand"
	"time"
)

type RPCDemoService struct{} // service name

// function to receive message from the server as a reply string
func (*RPCDemoService) GetServerMessage(args *Args, reply *string) error {
	*reply = "This is a message from RPC server!"
	return nil
}

// method to simulate long running call
func (*RPCDemoService) SomeLongRunningProcess(args *Args, reply *string) error {

	// sleep for a duration between 1 and 5 seconds
	time.Sleep(time.Second * (1 + time.Duration(rand.Intn(6))))
	*reply = "Success!"
	return nil
}
