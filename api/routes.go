package api

import (
	"net/http"
)

// RegisterRoutes sets up the API routes
func RegisterRoutes() {
	http.HandleFunc("/products", CreateProductHandler)
}
