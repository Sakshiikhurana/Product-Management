package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func BenchmarkGetProductNoCache(b *testing.B) {
	req, err := http.NewRequest("GET", "/products/1", nil)
	if err != nil {
		b.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetProductHandler)

	// Run benchmark test for `b.N` iterations
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		handler.ServeHTTP(rr, req)
	}
}
