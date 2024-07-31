package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var (
	brokerAddr = "localhost:1883"
	clientID   = "server_mqtt"
)

func checkError(t mqtt.Token) {
	if t.Error() != nil {
		log.Fatalf("Houve um erro: %v", t.Error())
	}
}

func main() {

	// Configuração do cliente MQTT
	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://" + brokerAddr)
	opts.SetClientID(clientID)

	client := mqtt.NewClient(opts)
	token := client.Connect()
	token.Wait()
	checkError(token)
	fmt.Println("Conectado ao broker MQTT com sucesso")

	token = client.Subscribe("http_req", 0, receiveHandler)
	token.Wait()
	checkError(token)
	fmt.Println("Assinado no tópico 'http_req'")

	select {}
}

var receiveHandler mqtt.MessageHandler = func(c mqtt.Client, m mqtt.Message) {
	//Deserializa o httpReq e manda para o processador de req
	req, _ := deserializeRequest(m.Payload())
	PublishFileContent(req, c)
}

func deserializeRequest(payload []byte) (*http.Request, error) {
	// Converte o payload para uma string
	payloadStr := string(payload)
	lines := strings.Split(payloadStr, "\n")

	// Cria uma nova requisição HTTP
	req := &http.Request{
		Method:     "GET",      // Valor default
		Proto:      "HTTP/1.1", // Valor default
		ProtoMajor: 1,          // Valor default
		ProtoMinor: 1,          // Valor default
		Header:     make(http.Header),
	}

	// Analisar a linha requerida
	if len(lines) > 0 {
		reqLine := lines[0]
		parts := strings.Fields(reqLine)
		if len(parts) >= 3 {
			req.Method = parts[0]
			req.RequestURI = parts[1] // Set the request URI
			req.Proto = parts[2]
		}
	}

	// Analisar os headers
	headersSection := false
	var headers []string
	for _, line := range lines[1:] {
		if line == "" {
			headersSection = true
			continue
		}
		if headersSection {
			headers = append(headers, line)
		}
	}

	// Adicionar os headers no request
	for _, header := range headers {
		headerParts := strings.SplitN(header, ":", 2)
		if len(headerParts) == 2 {
			key := strings.TrimSpace(headerParts[0])
			value := strings.TrimSpace(headerParts[1])
			req.Header.Add(key, value)
		}
	}

	// Adicionar o corpo
	if headersSection && len(lines) > 1 {
		body := strings.Join(lines[len(lines)-1:], "\n")
		req.Body = io.NopCloser(bytes.NewReader([]byte(body)))
	}

	return req, nil
}

func PublishFileContent(req *http.Request, client mqtt.Client) error {
	//Analisa se é do método GET
	if req.Method != "GET" {
		return nil
	}

	// Abre o arquivo
	filepath := "C:/Users/labou/OneDrive/Documentos/UFPE/CODING/codesGO/marmotas-3/MQTT/index.html"
	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("falha ao abrir o arquivo: %w", err)
	}
	defer file.Close()

	// Lê o conteúdo do arquivo
	bodyBytes, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("falha ao ler o arquivo: %w", err)
	}

	// Publica o conteúdo do arquivo no tópico MQTT
	token := client.Publish("index_html_body", 0, false, string(bodyBytes))
	token.Wait()
	if token.Error() != nil {
		return fmt.Errorf("falha ao publicar mensagem: %w", token.Error())
	}

	return nil
}
