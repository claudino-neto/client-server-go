package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Trying to connect with Server on \"localhost:8080\"...")
	client := http.Client{Timeout: time.Duration(1) * time.Second}
	// input >> get | post
	// url >> localhost:8080/...
	if err != nil {
		fmt.Printf("Error while dialing: %s\n", err.Error())
	}
	fmt.Println("Server connection estabilished")

}
