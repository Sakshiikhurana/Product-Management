// rabbitmq/main.go
package main

import (
    "fmt"
    "log"
    "github.com/rabbitmq/amqp091-go"
)

func main() {
    conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
    if err != nil {
        log.Fatalf("Failed to connect to RabbitMQ: %v", err)
    }
    defer conn.Close()

    ch, err := conn.Channel()
    if err != nil {
        log.Fatalf("Failed to open a channel: %v", err)
    }
    defer ch.Close()

    msgs, err := ch.Consume(
        "image_processing", // Queue name
        "",                 // Consumer name
        true,               // Auto-ack
        false,              // Exclusive
        false,              // No-local
        false,              // No-wait
        nil,                // Arguments
    )
    if err != nil {
        log.Fatalf("Failed to register a consumer: %v", err)
    }

    // Processing messages
    go func() {
        for d := range msgs {
            fmt.Printf("Received message: %s\n", d.Body)
            // Image processing logic here
        }
    }()

    fmt.Println("Waiting for messages. To exit press CTRL+C")
    select {} // Keep the program running
}
