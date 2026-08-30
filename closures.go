// Go supports anonymous functions which can form closures

package main

import "fmt"

// intSeq returns another function defined anonymously in its body
func intSeq() func() int {
	i := 0

	// This returned function closes over the variable i to form a closure
	return func() int {
		i++;
		return i
	}
}

func main() {
	// Variable stores the function result which captures its own i value
	nextInt := intSeq()

	// This i value is updated on each call to the function stored in nextInt
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// Confirms that the state (i) is unique to individual functions
	newInts := intSeq()
	fmt.Println(newInts()) // default (1) value for i; no updates from nextInt var
}
