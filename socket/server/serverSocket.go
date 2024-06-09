package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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
	var data map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println("Erro ao analisar o corpo da requisição:", err)
		http.Error(w, "Erro ao analisar o corpo da requisição", http.StatusBadRequest)
		return
	}

	val1, val1Exists := data["valor1"].(float64)
	val2, val2Exists := data["valor2"].(float64)
	operacao, operacaoExists := data["operacao"].(string)

	if !val1Exists || !val2Exists || !operacaoExists {
		fmt.Println("Valores ou operação não fornecidos pelo cliente")
		http.Error(w, "Valores ou operação não fornecidos pelo cliente", http.StatusBadRequest)
		return
	}

	// Realizando a operação
	var resultado float64
	switch operacao {
	case "+":
		resultado = val1 + val2
	case "-":
		resultado = val1 - val2
	case "*":
		resultado = val1 * val2
	case "/":
		if val2 == 0 {
			fmt.Println("Divisão por zero")
			http.Error(w, "Divisão por zero", http.StatusBadRequest)
			return
		}
		resultado = val1 / val2
	default:
		fmt.Println("Operação desconhecida:", operacao)
		http.Error(w, "Operação desconhecida", http.StatusBadRequest)
		return
	}

	// Respondendo ao cliente com o resultado
	fmt.Fprintf(w, "Resultado da operação %s entre %s e %s é %.2f", operacao, strconv.FormatFloat(val1, 'f', -1, 64), strconv.FormatFloat(val2, 'f', -1, 64), resultado)
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
