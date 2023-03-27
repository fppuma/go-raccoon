package main

import (
	"fmt"
	"sort"
)

/*
*
Function Literal as Callback function
sort.Slice(slice, callback)
*
*/
func main() {
	// Sort a slice of integers in descending order using a function literal
	nums := []int{5, 2, 8, 1, 9}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	fmt.Println(nums) // Output: [9 8 5 2 1]

}
