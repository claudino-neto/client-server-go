package server

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

func server() {
	httpRequest := new(impl.HTTP)
	server := rpc.NewServer()
	err := server.RegisterName("HTTP", httpRequest)

	ln, err := net.Listen("tcp", "localhost:1010")
	if err != nil {
		log.Fatal("Local Network not founded: ", err)
	}

	defer ln.Close()

	fmt.Println("Servidor está pronto ...")
	server.Accept(ln)
}

//TODO:
// - VERIFICAR PACKAGE
// = VER O DEFER
