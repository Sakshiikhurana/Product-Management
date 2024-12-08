package database

import (
	"testing"
	"os"
	"log"
)

func TestMain(m *testing.M) {
	// Initialize DB connection (make sure the test DB exists)
	err := InitDB()
	if err != nil {
		log.Fatal("Failed to initialize DB: ", err)
	}
	// Run tests
	exitCode := m.Run()
	// Perform cleanup if necessary (e.g., close DB connection)
	os.Exit(exitCode)
}

func TestCreateProduct(t *testing.T) {
	name := "Test Product"
	description := "A test product"
	price := "20.99"

	productID, err := CreateProduct(name, description, price)
	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}

	if productID <= 0 {
		t.Fatalf("Expected a valid product ID, but got %d", productID)
	}
}

func TestGetProductByID(t *testing.T) {
	// Assuming product ID 1 exists in your test DB
	productID := int64(1)

	name, description, price, err := GetProductByID(productID)
	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}

	if name == "" || description == "" || price == "" {
		t.Fatal("Expected non-empty product fields, but got empty values")
	}
}
