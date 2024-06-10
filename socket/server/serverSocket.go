package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "HOME PAGE!")
	// Respondendo ao cliente com o resultado
	fmt.Fprintf(w, "Resultado da requisção %s é HOME PAGE!", r.Method)
}

func handleReq(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Recebendo requisição do cliente...")

	// Lendo o corpo da requisição
	var data map[string]string
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println("Erro ao analisar o corpo da requisição:", err)
		http.Error(w, "Erro ao analisar o corpo da requisição", http.StatusBadRequest)
		return
	}

	link, linkExists := data["link"]

	if !linkExists {
		fmt.Println("Valores ou operação não fornecidos pelo cliente")
		http.Error(w, "Valores ou operação não fornecidos pelo cliente", http.StatusBadRequest)
		return
	}

	// Realizando a operação
	var reply string
	switch r.Method {
	case "GET":

		// Fazendo a solicitação GET
		response, err := http.Get(link)
		if err != nil {
			fmt.Println("Error to create solicitation:", err)
		}
		defer response.Body.Close() // Certifique-se de fechar o corpo da resposta

		// Verificando o código de status da resposta
		if response.StatusCode != http.StatusOK {
			fmt.Println("Error: status not OK", response.Status)
		}

		// Lendo o corpo da resposta
		body, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error to read reply body:", err)
		}

		// Imprimindo o corpo da resposta
		reply = string(body)

	case "HEAD":
		// Fazendo a solicitação HEAD
		response, err := http.Head(link)
		if err != nil {
			fmt.Println("Error to create solicitation:", err)
		}
		defer response.Body.Close() // Certifique-se de fechar o corpo da resposta

		// Verificando o código de status da resposta
		if response.StatusCode != http.StatusOK {
			fmt.Println("Error: status not OK", response.Status)
		}

		// Obtendo os cabeçalhos da resposta
		headers := response.Header
		reply = fmt.Sprintf("Headers: %v", headers)

	case "TRACE":
		req, err := http.NewRequest("TRACE", link, nil)
		if err != nil {
			fmt.Println("Error to create TRACE request:", err)
		}

		// Fazendo a solicitação TRACE
		client := &http.Client{}
		response, err := client.Do(req)
		if err != nil {
			fmt.Println("Error to create solicitation:", err)
		}
		defer response.Body.Close() // Certifique-se de fechar o corpo da resposta

		// Verificando o código de status da resposta
		if response.StatusCode != http.StatusOK {
			fmt.Println("Error: status not OK", response.Status)
		}

		// Lendo o corpo da resposta
		body, err := io.ReadAll(response.Body)
		reply = string(body)
		if err != nil {
			fmt.Println("Error to read reply body:", err)
		}
	default:
		fmt.Println("Requisição desconhecida:", reply)
		http.Error(w, "Requisição desconhecida", http.StatusBadRequest)
		return
	}
	// Respondendo ao cliente com o resultado
	fmt.Fprintf(w, "A requisição é %s", reply)
}

func setupRoutes() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/req", handleReq)
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
