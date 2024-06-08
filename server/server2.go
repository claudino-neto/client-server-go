package main

import (
	"fmt"
	"io"
	"net/http"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "HOME PAGE!")
}

func handleCalc(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Calculator Service")
	io.WriteString(w, "Calculator Service")
}

func setupRoutes() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/calculator", handleCalc)
}

func main() {
	setupRoutes()

	// Establishing TCP Connection
	fmt.Println("Server Running...")
	fmt.Println("Listening on \"localhost:8081\"")
	fmt.Println("Waiting for Client to connect...")

	err := http.ListenAndServe("localhost:8081", nil)

	if err != nil {
		fmt.Printf("Error while listening: %s\n", err.Error())
	}

	fmt.Println("Conexão encerrada!")

}
