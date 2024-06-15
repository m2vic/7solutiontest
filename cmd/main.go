package main

import (
	handler "gofortest/pkg/handlers"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/beef", handler.BeefHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
