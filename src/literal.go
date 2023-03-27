package main

import "fmt"

/*
*
In Go, an anonymous function is a function that doesn't have a name.
It's also sometimes referred to as a "function literal".
*
*/
func main() {
	// Declare and call an anonymous function that prints a message
	func(msg string) {
		fmt.Println(msg)
	}("Hello, world!")
}
