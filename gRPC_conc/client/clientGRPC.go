package main

import (
	"context"
	"fmt"
	pb "gRPC/gen"
	"os"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Args struct {
	A string
}

const GrpcPort = 1234

func ChecaErro(err error, msg string) {
	if err != nil {
		fmt.Println(msg, err)
		panic(err)
	}
}

func main() {

	// Creates a new file to hold the time durations
	file, err := os.Create("time.txt")
	if err != nil {
		fmt.Println("Failed to create file: ", err)
	}
	defer file.Close()

	//Cria as credenciais de transporte para garantir segurança
	opt := grpc.WithTransportCredentials(insecure.NewCredentials())
	endPoint := "localhost" + ":" + strconv.Itoa(GrpcPort)                     //define em que porta o client deve parar
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5000) //controlar o tempo máx que a função vai levar(1 seg no nosso caso)
	defer cancel()

	//cria a conexão nova
	conn, err := grpc.NewClient(endPoint, opt)
	ChecaErro(err, "Não foi possível se conectar ao servidor em "+endPoint)

	defer conn.Close()

	HTTPreq := pb.NewHTTPServiceClient(conn)

	for idx := 0; idx < 2; idx++ { // trocar o numero pra quantidade de requisições que você quer
		TempoInicio := time.Now()
		x, err := HTTPreq.GET(ctx, &pb.Request{Link: "http://cin.ufpe.br/~lab9"})
		ChecaErro(err, "Erro ao invocar a operação remota")
		fmt.Println(x.Body)
		TempoFim := time.Now()
		TempoTotal := TempoFim.Sub(TempoInicio)

		_, err = file.WriteString(TempoTotal.String() + "\n")
		if err != nil {
			fmt.Println("Failed to write to file: ", err)
		}
	}
}
