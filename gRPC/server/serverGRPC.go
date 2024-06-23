package main

import (
	"fmt"
	"net"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const GrpcPort = 50051

// Função para verificar e tratar erros
func ChecaErro(err error, msg string) {
	if err != nil {
		fmt.Println(msg, err)
		panic(err)
	}
}

func main() {
	endpoint := "localhost:" + strconv.Itoa(GrpcPort)
	conn, err := net.Listen("tcp", endpoint)
	ChecaErro(err, "Não foi possível criar o listener")
	server := grpc.NewServer()

	//gen.RegisterHTTPServer(server, &impl.HTTPproc{})
	reflection.Register(server)
	fmt.Println("Servidor pronto...")

	err = server.Serve(conn)
	ChecaErro(err, "Falha ao inciar servidor")

	//httpProc := impl.NewHTTPproc() // Use NewHTTPproc para inicializar corretamente
	// rpc.Register(httpProc)
	// rpc.HandleHTTP()
	// listener, err := net.Listen("tcp", ":1234")
	// if err != nil {
	// 	log.Fatal("Error while listening: ", err)
	// }
	// fmt.Println("Server aberto :-)")
	// http.Serve(listener, nil)
}
