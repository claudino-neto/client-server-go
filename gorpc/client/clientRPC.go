package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
	//"time"
)

type Args struct {
	A string
}

func main() {
	var reply string

	// Creates a new file to hold the time durations
	file, err := os.Create("time3.txt")
	if err != nil {
		fmt.Println("Failed to create file: ", err)
	}
	defer file.Close()

	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Error while dialing: ", err)
	}

	args := Args{A: "http://cin.ufpe.br/~lab9"}
	// Conectando com o server
	//for idx := 0; idx < 10000; idx++ { //trocar o numero pra quantidade de requisições que você quer
	//TempoInicio := time.Now()
	err = client.Call("HTTPproc.GET", args, &reply)
	if err != nil {
		fmt.Println("Error calling the server: ", err)
	}
	//TempoFim := time.Now()
	//TempoTotal := TempoFim.Sub(TempoInicio)

	// _, err = file.WriteString(TempoTotal.String() + "\n")
	// if err != nil {
	// 	fmt.Println("Failed to write to file: ", err)
	// }
	fmt.Println("Reply: %s", reply)
}

//}
