package main

import (
	"fmt"
	"log"

	"github.com/faizan2786/go-projects/RPC/internal/server"
	"github.com/faizan2786/go-projects/RPC/services"
)

func main() {

	client, err := server.NewClient() // Get the rpc client for the Demo service

	if err != nil {
		log.Fatalf("Client creation failed: %v", err)
	}
	defer client.Close() // close the connection when program ends

	// test demo service...

	args := services.Args{} // empty args
	var reply string        // variable to hold reply

	// Make the RPC call...
	// Call normally takes service_name.function_name, args and
	// the address of the variable that holds the reply.
	// Here we have no args in the RPC method therefore, we can pass the empty args struct.
	err = client.Call("RPCDemoService.GetServerMessage", &args, &reply)

	if err != nil {
		log.Fatal("Error while invoking RPC method on demo service!", err)
	}

	fmt.Println(reply)

	// test Arith service...

	// invoke Multiply method
	var mulResult int
	args = services.Args{A: 15, B: 6}
	err = client.Call("Arith.Multiply", &args, &mulResult)

	if err != nil {
		log.Fatal("Error while invoking Multiply method on Arith service!", err)
	}
	fmt.Printf("%d * %d = %d\n", args.A, args.B, mulResult)

	// invoke Divide method
	divResult := services.DivisionResult{}
	err = client.Call("Arith.Divide", args, &divResult)

	if err != nil {
		log.Fatal("Error while invoking Divide method on Arith service!", err)
	}

	fmt.Printf("%d / %d gives quotient = %d, remainder = %d\n", args.A, args.B, divResult.Quo, divResult.Rem)
}
