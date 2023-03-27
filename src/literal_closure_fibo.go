package main

import "fmt"

func fibonacci() func() int {
	prev1, prev2 := 0, 1
	return func() int {
		result := prev1
		prev1, prev2 = prev2, prev1+prev2
		return result
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 12; i++ {
		fmt.Println(f())
	}
}
