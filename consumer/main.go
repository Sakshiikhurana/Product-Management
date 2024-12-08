package main

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

type ImageMessage struct {
	URL string `json:"url"`
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		"image_processing", // Queue name
		"",                 // Consumer
		true,               // Auto-ack
		false,              // Exclusive
		false,              // No-local
		false,              // No-wait
		nil,                // Args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %s", err)
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			var msg ImageMessage
			json.Unmarshal(d.Body, &msg)
			log.Printf("Processing image: %s", msg.URL)
			// Add image compression logic here.
		}
	}()
	log.Println("Waiting for messages...")
	<-forever
}
