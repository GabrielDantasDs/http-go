package main

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	Name string
	URL  string
}

func Start() {
	http.Handle("/foo",  http.NotFoundHandler())

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
