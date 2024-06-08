package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Trying to connect with Server on \"localhost:8081\"...")
	client := http.Client{Timeout: time.Duration(1) * time.Second}
	// input >> get | post
	// url >> localhost:8080/...

	//Criando uma requisição HTTP GET para o servidor local
	req, err := http.NewRequest("GET", "http://localhost:8081/calculator", nil)
	if err != nil {
		fmt.Println("Erro ao criar requisição:", err)
		return
	}

	// Envia a requisição usando o cliente criado e recebe a resposta
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro ao enviar requisição:", err)
		return
	}
	defer resp.Body.Close()

	// Lê o corpo da resposta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler resposta:", err)
		return
	}

	// Processa a resposta
	fmt.Println("Resposta do servidor:", string(body))

}
