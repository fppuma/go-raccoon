package main

import (
	"fmt"
	"reflect"
)

// Define a parameterized function that returns a function literal
func isGreaterThan(threshold int) func(int) bool {
	return func(num int) bool {
		return num > threshold
	}
}

func main() {

	// Use the function to create a function literal that checks if a number is greater than 5
	checkGreaterThan5 := isGreaterThan(5)
	fmt.Println("Type of checkGreaterThan5:", reflect.TypeOf(checkGreaterThan5))
	// Output: Type of checkGreaterThan5: func(int) bool

	fmt.Println(checkGreaterThan5(3)) // Output: false
	fmt.Println(checkGreaterThan5(7)) // Output: true

}
