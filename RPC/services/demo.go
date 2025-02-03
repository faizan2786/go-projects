package services

import "time"

type RPCDemoService struct{} // service name

// function to receive message from the server as a reply string
func (*RPCDemoService) GetServerMessage(args *Args, reply *string) error {
	*reply = "This is a message from RPC server!"
	return nil
}

// method to simulate long running call
func (*RPCDemoService) SomeLongRunningProcess(args *Args, reply *string) error {
	time.Sleep(5 * time.Second)
	*reply = "Success!"
	return nil
}
