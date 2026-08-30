package main

import "fmt"

// Function that returns two integers
func twoVals() (int, int) {
	return 2, 3
}

func main() {
	// Retrieve both values and print
	x, y := twoVals()
	fmt.Println("x:", x)
	fmt.Println("y:", y)

	// Ignore first value, print second
	_, z := twoVals()
	fmt.Println("z:", z)
}
