package main

import (
	"log"
	"net/http"
	"os"
	"msa-pilot-devong-employee/router"
)

func main() {
	router := router.NewRouter()

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(":" + port, router))
}