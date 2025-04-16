package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		nodeName := os.Getenv("NODE_NAME") + os.Getenv("K_SERVICE")
		banner := "FROM: K8S NODE "
		if os.Getenv("K_SERVICE") != "" {
			banner = "FROM: CLOUDRUN "
		}

		resp := map[string]string{
			"host":   r.Host,
			"node":   nodeName,
			"source": r.RemoteAddr,
		}
		b, err := json.Marshal(resp)

		if err != nil {
			fmt.Fprintf(w, "failed to marshal json: %v\n", err)
			log.Printf("failed to marshal json: %v", err)
		} else {
			fmt.Fprint(w, banner+string(b)+"\n")
			log.Printf("reply with request: %s", string(b))
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "80" // Default port if PORT is not set
	}

	log.Printf("Server listening on port %s\n", port) // Log the port

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err) // Log fatal error and exit
	}
}
