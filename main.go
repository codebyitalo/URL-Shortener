package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Hello, %q", html.EscapeString(req.URL.Path))
	})

	http.HandleFunc("/hi", func(res http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(res, "Hi")
	})

	fmt.Printf("Servidor rodando na porta 8080\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
