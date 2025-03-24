package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		source := r.RemoteAddr
		host := r.Host
		scheme := "http://"
		if r.TLS != nil {
			scheme = "https://"
		}

		url := scheme + host

		response := fmt.Sprintf("Source: %s\tURL: %s\n", source, url)
		fmt.Fprintf(w, response)
		log.Printf(response)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if PORT is not set
	}

	log.Printf("Server listening on port %s\n", port) // Log the port

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err) // Log fatal error and exit
	}
}
