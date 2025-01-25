package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Client struct {
	Name string
	WebClient http.Client
}

func Request() {
	time.Sleep(1 * time.Second)
	client := Client{Name: "Gabriel", WebClient: *http.DefaultClient}

	response, error := client.WebClient.Get("http://127.0.0.1:8080/bar")

	if error != nil {
		log.Fatalf("Erro ao fazer requisição: %v", error)
	}

	defer response.Body.Close() // Fecha o corpo da resposta após a leitura

	// Lê o corpo da resposta
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Erro ao ler resposta: %v", err)
	}

	fmt.Println(string(body))
	fmt.Println(error)
}
