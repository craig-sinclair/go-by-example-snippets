package main

import "fmt"

func main() {
	var a[5] int
	a[0] = 100

	for i := range len(a) {
		fmt.Println(a[i])
	}

	b := [...]int{1,2,3,4,5}
	for j := range len(b) {
		fmt.Println(b[j])
	}

	fmt.Println(b)

}
