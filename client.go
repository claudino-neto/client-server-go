// package main

// import (
// 	"container/list"
// 	"fmt"
// 	"net/rpc"
// )

// func main() {
// 	var reply[] string
// 	var request string

//     fmt.Print("Qual é a sua requisição HTTP para o nosso web?(Digite HEAD, GET OU TRACE)")
//     fmt.Scan(&request)

// 	client, err := rpc.Dial("tcp", "localhost:1010")
// 	if err != nil {
// 		log.Fatal("dialing:", err)
// 	}

// 	defer func(client *rpc.Client){
// 		fmt.Println("oi")
// 	}

// 	args := shared.Args{request}
// 	err = client.Call("HTTP." + request,args, &reply)

// 	if err != nil{
// 		log.Fatal("service call error:", err)
// 	}
// 	else {
// 		fmt.println("HTTP" + request + "reply: " + reply)
// 	}
// }
// //TODO:
// // - MODIFICAR O DEFER(AINDA NÃO SEI O QUE COLOCAR)
// // - VERIFICAR O PACKAGE
