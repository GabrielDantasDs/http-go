package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Server struct {
	Name string
	URL  string
}

func Start() {
	http.Handle("/foo", http.NotFoundHandler())

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			requestBody, error := io.ReadAll(r.Body)

			if error != nil {
				log.Fatal("Error on handle request")
				return
			}

			defer r.Body.Close()

			fmt.Println(string(requestBody))
		}

		if r.Method == http.MethodGet {
			w.Write([]byte("Response from GET"))
		}
	})

	http.HandleFunc("/download", func(w http.ResponseWriter, r *http.Request) {
		filePath := "arquivo.txt"

		// Serve o arquivo para o cliente
		http.ServeFile(w, r, filePath)
	})

	http.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		user, password, _ := r.BasicAuth()

		fmt.Println("Auth credentials: ")
		fmt.Println(string(user))
		fmt.Println(string(password))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
