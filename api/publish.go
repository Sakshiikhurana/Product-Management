package api

import (
	"encoding/json"
	
	"github.com/rabbitmq/amqp091-go"
)

func publishToQueue(imageURLs []string) error {
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(
		"image_processing", // Queue name
		true,               // Durable
		false,              // Delete when unused
		false,              // Exclusive
		false,              // No-wait
		nil,                // Arguments
	)
	if err != nil {
		return err
	}

	for _, url := range imageURLs {
		body, _ := json.Marshal(map[string]string{"url": url})
		err = ch.Publish(
			"",         // Exchange
			"image_processing", // Routing key
			false,      // Mandatory
			false,      // Immediate
			amqp091.Publishing{
				ContentType: "application/json",
				Body:        body,
			},
		)
		if err != nil {
			return err
		}
	}
	return nil
}
