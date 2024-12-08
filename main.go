package main

import (
	"fmt"
	"log"
	"net/http"

	// Import your image processing package
	"product-management/api"
)

func main() {
	// Initialize routes for your API
	api.RegisterRoutes()

	// Start the web server on port 8080
	fmt.Println("Starting server on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
