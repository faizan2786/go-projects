package main

import (
	"fmt"
	"math/rand"
	"time"
)

// worker process that:
// receives requests on the request channel,
// reads the request off the queue
// process the request and put the response on output channel,
// then moves on to reading next available request on the channel
// i.e. This mechanism enables dynamic processing of requests by workers
func worker(id int, requests <-chan int, response chan<- string) {

	// keep processing requests as long as they are available
	for in := range requests {
		// simulate the request work by random delay
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		response <- fmt.Sprintf("Worker %d has finished a request %d", id, in)
	}
}

// spawns worker routines (i.e. distributes requests among workers)
// it uses fan-in pattern to collect the responses from workers to a single output channel
func loadBalancer(numWorkers int, requests <-chan int) <-chan string {
	out := make(chan string)
	for i := 0; i < numWorkers; i++ {
		go worker(i, requests, out)
	}
	return out
}

// client program that:
// creates the requests channel and passes it to the load balancer
// fires the requests onto the channel
// reads off the responses from the returned channel
func main() {

	const numWorkers int = 4
	const numRequests int = 10

	requests := make(chan int, 100) // buffered requests channel
	responseCh := loadBalancer(numWorkers, requests)

	// fire the requests (in the background)
	go func(ch chan<- int) {
		for i := 0; i < numRequests; i++ {
			requests <- i
		}
		close(requests) // close the request channel once done firing them all
	}(requests)

	// read the responses
	for i := 0; i < numRequests; i++ {
		fmt.Println(<-responseCh)
	}
}
