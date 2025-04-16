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
		r.ParseForm() // Best effort parsing form
		url := r.Form.Get("url")
		if url == "" {
			http.Error(w, "Missing 'url' parameter", http.StatusBadRequest)
			log.Print("Missing 'url' parameter")
			return
		}

		// Create a new request to the target URL
		// For simplicity, we are making a GET request here, regardless of the incoming method.
		// A more complete proxy would use r.Method and potentially r.Body.
		outReq, err := http.NewRequest("GET", url, nil) // Using GET as implied by original code
		if err != nil {
			http.Error(w, "Error creating request: "+err.Error(), http.StatusInternalServerError)
			log.Printf("Error creating request for %q: %v", url, err)
			return
		}

		// Copy headers from the incoming request (r) to the outgoing request (outReq)
		for name, values := range r.Header {
			// Don't copy hop-by-hop headers or the Host header.
			// The http client will set the Host header automatically based on the outReq.URL.
			// Connection-related headers are managed by the client.
			// Content-Length is managed by the client based on the body (which is nil here).
			if name == "Host" || name == "Connection" || name == "Keep-Alive" || name == "Proxy-Authenticate" || name == "Proxy-Authorization" || name == "Te" || name == "Trailers" || name == "Transfer-Encoding" || name == "Upgrade" || name == "Content-Length" {
				continue
			}
			for _, value := range values {
				outReq.Header.Add(name, value)
			}
		}
		// Optionally, add or modify specific headers if needed, e.g.:
		// outReq.Header.Set("X-Forwarded-For", r.RemoteAddr)

		// Execute the outgoing request
		client := &http.Client{}
		resp, err := client.Do(outReq)
		if err != nil {
			http.Error(w, "Error fetching URL: "+err.Error(), http.StatusBadGateway)
			log.Printf("Error fetching %q: %v", url, err)
			return
		}
		defer resp.Body.Close()

		// --- Response Handling (similar to original code) ---

		// // OPTIONAL: Copy headers from the upstream response (resp) to the client response (w)
		// for name, values := range resp.Header {
		// 	for _, value := range values {
		// 		w.Header().Add(name, value)
		// 	}
		// }

		// // OPTIONAL: Set the status code from the upstream response
		// w.WriteHeader(resp.StatusCode)

		// Copy the body from the upstream response to the client response
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			// Handle potential errors during reading the body (e.g., connection closed prematurely)
			// We might have already started writing headers/status, so sending an HTTP error might not work well.
			log.Printf("Error reading response body from %q: %v", url, err)
			// Attempt to inform the client if possible, though headers/status might be sent.
			fmt.Fprintln(w, "\nError reading upstream response body.")
		} else {
			// If headers/status were not copied/set above, they will be set implicitly here
			// when writing the body (defaulting to 200 OK).
			// To properly proxy status and headers, uncomment the sections above.
			fmt.Fprintln(w, string(bodyBytes))
			log.Printf("Success proxy request to %q. code %d, resp: %v", url, resp.StatusCode, string(bodyBytes))
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
