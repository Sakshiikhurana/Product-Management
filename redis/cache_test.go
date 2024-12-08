package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/go-redis/redis/v8"
	"context"
	"fmt"
)

var ctx = context.Background()

func TestGetProductWithCache(t *testing.T) {
	// Setup Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// Make a request to the product API endpoint
	req, err := http.NewRequest("GET", "/products/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Record the response
	rr := httptest.NewRecorder()

	// Assuming you have a function to serve the API
	handler := http.HandlerFunc(GetProductHandler)
	handler.ServeHTTP(rr, req)

	// First request should result in a cache miss
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected status code %v, got %v", http.StatusOK, status)
	}

	// Check Redis cache for the product data
	cacheData, err := rdb.Get(ctx, "product:1").Result()
	if err != nil {
		t.Fatal(err)
	}

	if cacheData == "" {
		t.Errorf("Expected data in cache for product 1")
	}

	// Now, make the same request again, which should result in a cache hit
	req, err = http.NewRequest("GET", "/products/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	// Check that the second request served the data from the cache
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected status code %v, got %v", http.StatusOK, status)
	}
}
