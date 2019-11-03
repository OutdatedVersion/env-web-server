package main

import (
	"os"
	"net/http"
	"log"
	"fmt"
)

var (
	body = os.Getenv("BODY")
)

func handle(writer http.ResponseWriter, req *http.Request) {
	if body == "" {
		writer.WriteHeader(204)
	} else {
		fmt.Fprintln(writer, body)
	}
}

func main() {
	http.HandleFunc("/", handle)
	log.Printf("body: '%s'\n", body)

	http.ListenAndServe(":2000", http.DefaultServeMux)
}