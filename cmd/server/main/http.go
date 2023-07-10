package main

import (
	bootstrap "art-item/internal/bootstrap"
	"log"
)

func main() {
	err := bootstrap.InitHTTPServer()

	if err != nil {
		log.Fatalf("Failed to initialize HTTP server: %v", err)
	}
}
