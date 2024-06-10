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

	var link string
	var requestHttp string
	var keepGoing string

	for {
		fmt.Println("Qual é o link do site web que deverá processar o pedido http?(EX: http://cin.ufpe.br/~lab9)")
		fmt.Scanln(&link)

		fmt.Println("Digite a requisição HTTP que você quer enviar para o web. Escolha entre GET, HEAD ou TRACE")
		fmt.Scanln(&requestHttp)

		body, _ := json.Marshal(map[string]string{
			"link": "http://cin.ufpe.br/~lab9",
		})
		payload := bytes.NewBuffer(body)
		// Criando uma requisição HTTP para o servidor local
		req, err := http.NewRequest("GET", "http://localhost:8081/req", payload)
		if err != nil {
			fmt.Println("Erro ao criar requisição:", err)
			return
		}

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
		//Imprimindo a resposta do servidor
		fmt.Println("Resposta do servidor:", string(respBody))

		fmt.Println("Quer fazer mais uma requisição? (S/N)")
		fmt.Scanln(&keepGoing)
		if keepGoing == "N" {
			break
		}

	}

}
