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

	var valor1, valor2 int
	var operacao string

	fmt.Print("Digite o primeiro valor: ")
	fmt.Scanln(&valor1)

	fmt.Print("Digite o segundo valor: ")
	fmt.Scanln(&valor2)

	fmt.Print("Digite a operação (+, -, *, /): ")
	fmt.Scanln(&operacao)

	// valor1 := 2
	// valor2 := 5
	// operacao := '/'

	body, _ := json.Marshal(map[string]interface{}{
		"valor1":   valor1,
		"valor2":   valor2,
		"operacao": string(operacao),
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
