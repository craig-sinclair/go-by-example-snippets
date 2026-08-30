package main

import (
	"fmt"
)

func main() {
	a := make([] int, 3)
	for i := range 3 {
		a[i] = i
	}

	fmt.Println("a: ", a)

	a = append(a, 3)
	fmt.Println("a after append: ", a)

	c := make([] int, len(a))
	copy(c, a)
	fmt.Println("copy: ", c)

	slc := c[1:3]
	fmt.Println("[1:3] slice: ", slc)
}
