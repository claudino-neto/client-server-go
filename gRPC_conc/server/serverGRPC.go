package main

import (
	"fmt"
	pb "gRPC/gen"
	"gRPC/impl"
	"net"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const GrpcPort = 1234

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

	pb.RegisterHTTPServiceServer(server, &impl.HTTPproc{})
	reflection.Register(server)
	fmt.Println("Servidor pronto...")

	err = server.Serve(conn)
	ChecaErro(err, "Falha ao inciar servidor")
}
