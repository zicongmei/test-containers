package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func queryURL(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error querying %s: %v", url, err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response from %s: %v", url, err)
		return
	}

	log.Printf("Response: %s", body)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("missing the URL")
		os.Exit(1)
	}

	url := os.Args[1] // Get the URL from the first command-line argument

	for {
		queryURL(url)
		time.Sleep(1 * time.Second)
	}
}
