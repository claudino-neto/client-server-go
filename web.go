package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//HTTP GET

type HTTP struct{}

// HTTP GET
func (t *HTTP) GET(args *shared.Args, reply *string) error {
	url := "https://www.cin.ufpe.br/~lab9/"

	// Fazendo a solicitação GET
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error to create solicitation:", err)
		return
	}
	defer response.Body.Close() // Certifique-se de fechar o corpo da resposta

	// Verificando o código de status da resposta
	if response.StatusCode != http.StatusOK {
		fmt.Println("Error: status not OK", response.Status)
		return
	}

	// Lendo o corpo da resposta
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error to read reply body:", err)
		return
	}

	// Imprimindo o corpo da resposta
	fmt.Println("Reply:", string(body))
}

// HTTP HEAD

func (t *HTTP) HEAD(args *shared.Args, reply *int) error {
	*reply = make([]int, 0)
	for i := args.X; i <= args.Y; i++ {
		*reply = append(*reply, i)
	}
	return nil
}

// HTTP TRACE

func (t *HTTP) TRACE(args *shared.Args, reply *int) error {
	*reply = make([]string, 0)
	for i := args.X; i <= args.Y; i++ {
		*reply = append(*reply, i)
	}
	return nil
}

//TODO:
// -TEM QUE ENTENDER A QUESTAO DO PACKAGE IMPL OU MAIN OU SERVER OU CLIENT(ESTOU UM POUCO CONFUSO EM REALAÇÃO A ISSO)
// - PESQUISAR COMO FAZER O HTTP TRACE E HEAD EM GO
