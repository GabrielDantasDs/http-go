package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type Client struct {
	Name            string
	WebClient       http.Client
	AuthCredentials AuthCredentials
}

type AuthCredentials struct {
	User     string
	Password string
}

func RequestPost() {
	time.Sleep(1 * time.Second)
	client := Client{Name: "Gabriel", WebClient: *http.DefaultClient}
	requestBody := strings.NewReader("Hello, Reader!")
	response, error := client.WebClient.Post("http://127.0.0.1:8080/bar", "text", requestBody)

	if error != nil {
		log.Fatalf("Request error: %v", error)
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

func RequestGet() {
	time.Sleep(1 * time.Second)
	client := Client{Name: "Gabriel", WebClient: *http.DefaultClient}

	response, error := client.WebClient.Get("http://127.0.0.1:8080/bar")

	if error != nil {
		log.Fatalf("Request error: %v", error)
	}

	defer response.Body.Close() // Fecha o corpo da resposta após a leitura

	// Lê o corpo da resposta
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Error on read response body: %v", err)
	}

	fmt.Println(string(body))
	fmt.Println(error)
}

func RequestFile() {
	time.Sleep(1 * time.Second)
	client := Client{Name: "Gabriel", WebClient: *http.DefaultClient}

	response, error := client.WebClient.Get("http://127.0.0.1:8080/download")

	if error != nil {
		log.Fatalf("Request error: %v", error)
	}

	defer response.Body.Close() // Fecha o corpo da resposta após a leitura

	// Lê o corpo da resposta
	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatalf("Error on create file")
	}

	file, err := os.Create("arquivo_baixado.txt")

	if err != nil {
		log.Fatalf("Error on create file")
	}

	file.Write([]byte(body))

	file.Close()

	if err != nil {
		log.Fatalf("Erro ao ler resposta: %v", err)
	}

	fmt.Println(string(body))
	fmt.Println(error)
}

func RequestAuth() {
	time.Sleep(1 * time.Second)
	authCredentials := AuthCredentials{User: "Gabriel", Password: "123456"}
	encodeCredentials := base64.StdEncoding.EncodeToString([]byte(authCredentials.User + ":" + authCredentials.Password))

	client := Client{Name: "Gabriel", WebClient: *http.DefaultClient, AuthCredentials: authCredentials}
	var buf bytes.Buffer

	json.NewEncoder(&buf).Encode(client.AuthCredentials)

	// response, error := client.WebClient.Post("http://127.0.0.1:8080/auth", "application/json", &buf)

	request, error := http.NewRequest("POST", "http://127.0.0.1:8080/auth", &buf)

	if error != nil {
		log.Fatalf("Request error: %v", error)
	}

	request.Header.Add("Authorization", "Basic "+encodeCredentials)
	response, error := client.WebClient.Do(request)

	if error != nil {
		log.Fatalf("Request error: %v", error)
	}

	body, error := io.ReadAll(response.Body)

	if error != nil {
		log.Fatalf("Request error: %v", error)
	}

	fmt.Println(string(body))
}
