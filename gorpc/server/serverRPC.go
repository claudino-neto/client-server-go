package main

import (
	"gorpc/impl"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	calculator := new(impl.Calculator)
	rpc.RegisterName("Calculator", calculator)
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Error while listening: ", err)
	}
	http.Serve(listener, nil)
}
