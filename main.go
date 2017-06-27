package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	router := NewRouter()

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	log.Fatal(http.ListenAndServe(":" + port, router))
}