package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/faizan2786/go-projects/RPC/internal/server"
	"github.com/faizan2786/go-projects/RPC/services"
)

const timeout int = 3 // timeout in seconds

func main() {

	args := services.Args{}
	var res string

	// define a timeout context to handle request timeouts
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(timeout))
	defer cancel()

	client, err := server.NewClient()
	if err != nil {
		log.Fatal("fail to connect to the server")
	}

	rpcDone := make(chan bool) // channel to check if rpc call was successful
	fmt.Println("Making RPC call to SomeLongRunningProcess...")
	go func() {
		client.Call("RPCDemoService.SomeLongRunningProcess", &args, &res)
		rpcDone <- true
	}()

	select {
	case <-ctx.Done():
		fmt.Println("Timeout reached. Shutting down client")
	case <-rpcDone:
		fmt.Println("RPC call finished")
	}
	client.Close()
}
