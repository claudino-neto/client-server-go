package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func checkError(t mqtt.Token) {
	if t.Error() != nil {
		log.Fatalf("Houve um erro: %v", t.Error())
	}
}

func main() {
	//Criar o arquivo de tempo
	file, err := os.Create("time.txt")
	if err != nil {
		fmt.Println("Failed to create file: ", err)
	}
	defer file.Close()

	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://localhost:1883")
	opts.SetClientID("client_mqtt")

	// Cria uma nova instância de cliente MQTT
	client := mqtt.NewClient(opts)

	// Conecta ao broker MQTT
	token := client.Connect()
	token.Wait()
	checkError(token)
	fmt.Println("Conexão estabelecida com sucesso")

	defer client.Disconnect(250)

	for idx := 0; idx < 10000; idx++ {
		TempoInicio := time.Now()

		//torna o client um subscriber para receber a resposta do server
		token = client.Subscribe("index_html_body", 0, receiveHandler)
		token.Wait()
		checkError(token)

		// Cria a requisição HTTP
		req := &http.Request{
			Method: "GET",
			Header: make(http.Header),
		}

		// Serialização da requisição HTTP
		payload := serializeRequest(req)

		// Publica a mensagem no tópico "http_req"
		token = client.Publish("http_req", 0, false, payload)
		token.Wait()
		checkError(token)
		fmt.Println("Mensagem publicada com sucesso")

		TempoFim := time.Now()
		TempoTotal := TempoFim.Sub(TempoInicio)
		_, err = file.WriteString(TempoTotal.String() + "\n")
	}
	//Mantém a conexão aberta
	select {}
}

var receiveHandler mqtt.MessageHandler = func(c mqtt.Client, m mqtt.Message) {
	//fmt.Println(string(m.Payload()))
}

func serializeRequest(req *http.Request) []byte {
	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("Method: %s\n", req.Method))
	buffer.WriteString(fmt.Sprintf("Proto: %s\n", req.Proto))
	buffer.WriteString(fmt.Sprintf("ProtoMajor: %d\n", req.ProtoMajor))
	buffer.WriteString(fmt.Sprintf("ProtoMinor: %d\n", req.ProtoMinor))

	// Adiciona os headers
	buffer.WriteString("Headers:\n")
	for key, values := range req.Header {
		for _, value := range values {
			buffer.WriteString(fmt.Sprintf("%s: %s\n", key, value))
		}
	}

	// Adiciona o corpo (se existir)
	if req.Body != nil {
		bodyBytes, err := io.ReadAll(req.Body)
		if err != nil {
			log.Fatalf("Erro ao ler o corpo da requisição: %v", err)
		}
		buffer.WriteString("\nBody:\n")
		buffer.Write(bodyBytes)
	}

	return buffer.Bytes()
}
