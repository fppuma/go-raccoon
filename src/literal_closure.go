package main

import (
	"fmt"
	"reflect"
)

// Define a closure that remembers a counter variable
// - Function Name: makeCounter
// - Returns: A function
func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {

	// Use the closure to create a counter function
	counter01 := makeCounter()
	fmt.Println("Type of counter01:", reflect.TypeOf(counter01))
	fmt.Println(counter01()) // Output: 1
	fmt.Println(counter01()) // Output: 2
	fmt.Println(counter01()) // Output: 3

	fmt.Println("..creating another counter.")

	// Use the closure to create a counter function
	counter02 := makeCounter()
	fmt.Println("Type of counter02:", reflect.TypeOf(counter02))
	fmt.Println(counter02()) // Output: 1
	fmt.Println(counter02()) // Output: 2

}
