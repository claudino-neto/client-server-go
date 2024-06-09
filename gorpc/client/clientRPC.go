package main

import (
	"log"
	"net/rpc"
)

type Args struct {
	A float32
	B float32
}

func main() {
	var reply float64
	args := Args{A: 15, B: 29}

	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Error while dialing: ", err)
	}

	// Requesting Calculator
	err = client.Call("Calculator.Div", args, &reply)
	if err != nil {
		log.Fatal("Error calling the server: ", err)
	}
	log.Printf("Reply: %.2f", reply)

}
