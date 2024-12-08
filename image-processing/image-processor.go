package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/streadway/amqp"
)

// StartProcessing handles the connection to RabbitMQ and message consumption
func StartProcessing() {
	// Establish a connection to RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	// Create a channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	// Consume messages from the "image_processing" queue
	msgs, err := ch.Consume(
		"image_processing", // queue
		"",                 // consumer name
		true,               // auto-acknowledge
		false,              // exclusive
		false,              // no-local
		false,              // no-wait
		nil,                // arguments
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	// Start consuming messages from the queue
	forever := make(chan bool)

	go func() {
		for d := range msgs {
			var data map[string]string
			json.Unmarshal(d.Body, &data)

			url := data["url"]
			log.Printf("Received image URL: %s", url)

			err := processImage(url)
			if err != nil {
				log.Printf("Failed to process image: %v", err)
				continue
			}

			log.Printf("Successfully processed image: %s", url)
		}
	}()

	log.Println("Waiting for messages... Press Ctrl+C to exit")
	<-forever
}

// processImage simulates image processing
func processImage(imageURL string) error {
	// Simulate image processing logic
	log.Printf("Processing image: %s", imageURL)
	time.Sleep(2 * time.Second)
	if imageURL == "fail" {
		return log.Output(1, "Processing failed for URL: "+imageURL)
	}
	return nil
}
