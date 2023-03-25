package main

import "fmt"

func main() {
	// Create a new goroutine that prints a message
	go func() {
		fmt.Println("Hello, world!")
	}()

	// Wait for a key press before exiting
	fmt.Scanln()
}
