package main

import "fmt"

func main() {
	// Declare a function literal that takes two integers and returns their sum
	add := func(a, b int) int {
		return a + b
	}
	// We assign this function literal to a variable add,
	// which can then be used as a regular function.

	// Call the function literal and print the result
	sum := add(3, 4)
	fmt.Println(sum) // Output: 7

	fmt.Println(add(3, -4)) // Output: -1
	fmt.Println(add(4, -4)) // Output: 0
}
