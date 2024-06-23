package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Args struct {
	A string
}

const GrpcPort = 50051

func ChecaErro(err error, msg string) {
	if err != nil {
		fmt.Println(msg, err)
		panic(err)
	}
}

func main() {

	var idx int32

	opt := grpc.WithTransportCredentials(insecure.NewCredentials())
	endPoint := "localhost" + ":" + strconv.Itoa(GrpcPort)
	conn, err := grpc.Dial(endPoint, opt)
	ChecaErro(err, "Não foi possível se conectar ao servidor em "+endPoint)

	defer conn.Close()

	calc := gen.NewCalculadoraClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	for idx = 0; idx < shared.SampleSize; idx++ {
		x, err := calc.Add(ctx, &gen.Request{P1: 1, P2: 2})
		ChecaErro(err, "Erro ao invocar a operação remota")
		fmt.Println(x.N)
	}
}
