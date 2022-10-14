package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Go Bananas")
	}

	http.HandleFunc("/hello", helloHandler)

	log.Println("app listing on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
