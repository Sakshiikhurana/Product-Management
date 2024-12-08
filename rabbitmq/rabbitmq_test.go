package rabbitmq

import (
	"testing"
	"time"
	"log"
)

func TestPublishMessage(t *testing.T) {
	conn, ch, err := Connect()
	if err != nil {
		t.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()
	defer ch.Close()

	message := "Test Message"
	err = PublishMessage(ch, message)
	if err != nil {
		t.Fatalf("Failed to publish message: %v", err)
	}
}

func TestConsumeMessages(t *testing.T) {
	conn, ch, err := Connect()
	if err != nil {
		t.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()
	defer ch.Close()

	done := make(chan bool)
	go ConsumeMessages(ch, done)

	// Wait for a message to be consumed
	select {
	case <-done:
		// Message was consumed successfully
	case <-time.After(5 * time.Second):
		t.Fatal("Failed to receive a message")
	}
}
