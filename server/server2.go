package main

import (
	"fmt"
	"net/http"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "HOME PAGE!")
}

func handleCalc(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Calculator Service")
}

func setupRoutes() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/calculator", handleCalc)
}

func main() {
	// Establishing TCP Connection
	fmt.Println("Server Running...")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Printf("Error while listening: %s\n", err.Error())
	}
	fmt.Println("Listening on \"localhost:8080\"")
	fmt.Println("Waiting for Client to connect...")

	setupRoutes()
}
