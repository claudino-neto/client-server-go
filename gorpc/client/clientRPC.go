package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)

type Args struct {
	A string
}

func main() {
	var reply string
	var link string

	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Error while dialing: ", err)
	}

	args := Args{A: link}
	// Conectando com o server
	for idx := 0; idx < 1; idx++ { //trocar o numero pra quantidade de requisições que você quer
		TempoInicio := time.Now()
		err = client.Call("HTTPproc.GET", args, &reply)
		TempoFim := time.Now()
		TempoTotal := TempoFim.Sub(TempoInicio)
		fmt.Println(TempoTotal)
	}

	if err != nil {
		log.Fatal("Error calling the server: ", err)
	}

	log.Printf("Reply: %s", reply)

}
