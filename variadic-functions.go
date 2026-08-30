package main

import "fmt"

// Function that accepts an arbitrary number of int values
func sum(nums ...int) {
	total := 0

	for _, num := range nums {
		total += num
	}

	fmt.Println("Total: ", total)
}

func main() {
	// Use variadic function with individual arguments
	sum(1,2)
	sum(3,4,5)

	// Use multiple arguments in a slice to apply the variadic function
	nums := []int{1,2,3,4,5}
	sum(nums...)
}
