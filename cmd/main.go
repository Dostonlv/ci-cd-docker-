package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Dostonlv/ci-cd-docker/handlers"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", handlers.HelloHandler)
	http.HandleFunc("/health", handlers.HealthCheckHandler)

	log.Printf("Server starting on port %s", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
