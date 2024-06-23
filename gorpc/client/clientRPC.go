package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Args struct {
	A string
}

func main() {
	var reply string
	var link string
	var keepGoing string

	for {
		fmt.Println("Qual é o link do site web que deverá processar o pedido http?(EX: http://cin.ufpe.br/~lab9)")
		fmt.Scanln(&link)

		args := Args{A: link}

		// Conectando com o server
		client, err := rpc.DialHTTP("tcp", "localhost:1234")
		if err != nil {
			log.Fatal("Error while dialing: ", err)
		}

		err = client.Call("HTTPproc.GET", args, &reply)
		if err != nil {
			log.Fatal("Error calling the server: ", err)
		}

		log.Printf("Reply: %s", reply)
		fmt.Println("Quer fazer mais uma requisição? (S/N)")
		fmt.Scanln(&keepGoing)
		if keepGoing == "N" {
			break
		}
	}
}
