package main

import (
	"fmt"
)

func main() {
	// Create a new channel of type string
	message := make(chan string)

	// Spawn a new goroutine to send a message on the channel
	go func() {
		message <- "Hello, world!"
	}()

	// Wait for a message to be received on the channel
	answer := <-message
	fmt.Println("..answer received from goroutine.")

	// Print the message
	fmt.Printf("msg: %s\n", answer)
	fmt.Println("..end of main.")
}
