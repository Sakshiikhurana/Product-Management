package api

import (
	"encoding/json"
	
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func init() {
	// Configure logger to output structured JSON
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.InfoLevel)
}

// CreateProductHandler handles the creation of products
func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	logger.WithFields(logrus.Fields{
		"method": r.Method,
		"url":    r.URL.Path,
	}).Info("Received Create Product request")

	// Simulate product creation
	time.Sleep(1 * time.Second)

	// Simulate response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Product created"})

	logger.WithFields(logrus.Fields{
		"status":  http.StatusCreated,
		"elapsed": time.Since(start),
	}).Info("Product created successfully")
}
