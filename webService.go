package server.service

import (
	"errors"
	"fmt"
	"net/http"
)

func server() {
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}

}
