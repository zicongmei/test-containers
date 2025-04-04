package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		url := r.Form.Get("url")
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprint(w, err.Error()+"\n")
			log.Print(err.Error())
			return
		}
		defer resp.Body.Close()

		// Copy headers from the upstream response to the client response
		for name, values := range resp.Header {
			for _, value := range values {
				w.Header().Add(name, value)
			}
		}

		// Set the status code from the upstream response
		w.WriteHeader(resp.StatusCode)

		// Copy the body from the upstream response to the client response
		_, err = io.Copy(w, resp.Body)
		if err != nil {
			// Log the error, but we might have already started writing the response
			log.Print("Error copying response body:" + err.Error())
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
