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
	var requestHttp string
	var keepGoing string

	for {
		fmt.Println("Qual é o link do site web que deverá processar o pedido http?(EX: http://cin.ufpe.br/~lab9)")
		fmt.Scanln(&link)
		fmt.Println("Digite a requisição HTTP que você quer enviar para o web. Escolha entre GET, HEAD ou TRACE")
		fmt.Scanln(&requestHttp)

		args := Args{A: string(link)}

		// Requesting Calculator
		client, err := rpc.DialHTTP("tcp", "localhost:1234")
		if err != nil {
			log.Fatal("Error while dialing: ", err)
		}
		err = client.Call("HTTPproc.GET", args, &reply)
		if err != nil {
			log.Fatal("Error calling the server: ", err)
		}
		log.Println("Reply: %s", reply)

		fmt.Println("Quer fazer mais uma requisição? (S/N)")
		fmt.Scanln(&keepGoing)
		if keepGoing == "N" {
			break
		}
	}

}
