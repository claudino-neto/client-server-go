package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "HOME PAGE!")
}

func handleCalc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Recebendo requisição do cliente...")

	// Verificando se o método da requisição é POST
	if r.Method != http.MethodPost {
		http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
		return
	}

	// Lendo o corpo da requisição
	var data map[string]int
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println("Erro ao analisar o corpo da requisição:", err)
		http.Error(w, "Erro ao analisar o corpo da requisição", http.StatusBadRequest)
		return
	}

	val1, val1Exists := data["valor1"]
	val2, val2Exists := data["valor2"]

	if !val1Exists || !val2Exists {
		fmt.Println("Valores não fornecidos pelo cliente")
		http.Error(w, "Valores não fornecidos pelo cliente", http.StatusBadRequest)
		return
	}

	// Realizando a soma
	resultado := val1 + val2

	// Respondendo ao cliente com o resultado
	fmt.Fprintf(w, "Soma de %d e %d é %d", val1, val2, resultado)
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
