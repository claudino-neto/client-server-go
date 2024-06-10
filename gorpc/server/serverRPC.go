package main

import (
	"fmt"
	"gorpc/impl"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	httpProc := new(impl.HTTP)
	rpc.RegisterName("HTTPproc", httpProc)
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Error while listening: ", err)
	}
	fmt.Println("Server aberto :-)")
	http.Serve(listener, nil)

}
