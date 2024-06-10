package impl

import (
	"fmt"
	"io"
	"net/http"
)

type Args struct {
	A string
}

type HTTP struct{}

func (c *HTTP) GET(args *Args, reply *string) error {
	url := args.A

	// Fazendo a solicitação GET
	response, err := http.Get(url)
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
	*reply = string(body)
	return nil
}

func (c *HTTP) HEAD(args *Args, reply *string) error {
	url := args.A

	// Fazendo a solicitação HEAD
	response, err := http.Head(url)
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
	*reply = fmt.Sprintf("Headers: %v", headers)
	return nil
}

func (c *HTTP) TRACE(args *Args, reply *string) error {
	url := args.A

	// Criando a solicitação TRACE
	req, err := http.NewRequest("TRACE", url, nil)
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
	if err != nil {
		fmt.Println("Error to read reply body:", err)
	}

	// Imprimindo o corpo da resposta
	*reply = string(body)
	return nil
}
