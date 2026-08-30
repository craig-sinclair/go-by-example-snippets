package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

// Multiple parameters with the same type (int) and int return type specified
func plusThree(a, b, c int) int {
	return a + b + c
}

func printGreeting() {
	fmt.Println("hello world")
}

func main() {
	plusRes := plus(3, 4)
	fmt.Println("Plus result:", plusRes)

	a, b, c := 1, 2, 3
	resThree := plusThree(a, b, c)
	fmt.Println("Plus three result:", resThree)

	printGreeting()
	printGreeting()
}
