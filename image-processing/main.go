package main

import (
	"fmt"
     "product-management/image-processing"
)

func main() {
	// Print message to show that the service is running
	fmt.Println("Starting the Image Processing Service...")

	// Call StartProcessing function to handle RabbitMQ logic
	StartProcessing()
}
