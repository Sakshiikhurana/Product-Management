package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateProductHandler(t *testing.T) {
	// Simulate a request to the API endpoint
	req, err := http.NewRequest("POST", "/products", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Use a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Create a handler function (the function you're testing)
	handler := http.HandlerFunc(CreateProductHandler)

	// Call the handler with the request and recorder
	handler.ServeHTTP(rr, req)

	// Check if the response status code is what you expect
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("expected status code %v, got %v", http.StatusCreated, status)
	}
}
