package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/faizan2786/go-projects/RPC/services"
)

const hostName string = "" // left blank to serve on localhost
const portNum string = "1298"
const address = hostName + ":" + portNum

// function to start RPC server and continuously listen for http requests
func StartServer() {

	// create and register the services
	service1 := new(services.RPCDemoService)
	rpc.Register(service1) // registers with Go's default RPC server

	service2 := new(services.Arith)
	rpc.Register(service2)

	rpc.HandleHTTP()

	// listen for request on above port
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("Error while start a listener:\n", err)
	}

	// Handle server shutdowns gracefully using cancellable context and OS signals

	ctx, cancel := context.WithCancel(context.Background()) // a context that can be cancelled explicitly using cancelFunc
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM) // notify on Ctrl+C and process termination

	fmt.Println("A server is running to accept RPC calls...")

	// start a go routine to handle terminations
	go func() {
		<-signalChan // wait for a (termination) signal
		fmt.Println("Termination signal received. Shutting down server...")
		cancel() // cancel the context (sending cancel signal to all go routines)
		listener.Close()

		// give server some time to finish current request
		time.Sleep(time.Second)

		// exit gracefully
		os.Exit(0)
	}()

	// start the (default) server in a separate go routine
	go func() {
		http.Serve(listener, nil)
	}()

	// wait for context cancellation (i.e. sever shutdown) event
	<-ctx.Done()
	fmt.Println("Server shutdown complete.")
}

// wrapper function to instantiate the client
func NewClient() (*rpc.Client, error) {

	// create an rpc client and connect it to the server
	client, err := rpc.DialHTTP("tcp", address)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the server: %w", err)
	}
	return client, nil
}
