package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Trying to connect with Server on \"localhost:8081\"...")

	valor1 := 2
	valor2 := 5

	body, _ := json.Marshal(map[string]int{
		"valor1": valor1,
		"valor2": valor2,
	})
	payload := bytes.NewBuffer(body)

	// Criando uma requisição HTTP POST para o servidor local
	req, err := http.NewRequest("POST", "http://localhost:8081/calculator", payload)
	if err != nil {
		fmt.Println("Erro ao criar requisição:", err)
		return
	}

	// Configurando o cabeçalho da requisição
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	// Enviando a requisição usando o cliente criado e recebendo a resposta
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro ao enviar requisição:", err)
		return
	}
	defer resp.Body.Close()

	// Lendo o corpo da resposta
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler resposta:", err)
		return
	}

	// Imprimindo a resposta do servidor
	fmt.Println("Resposta do servidor:", string(respBody))

}
