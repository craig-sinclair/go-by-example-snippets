package main

import "fmt"

func main() {
	for i := 0; i<5; i++ {
		fmt.Println(i)
	}

	j := 0
	for j<5 {
		fmt.Println(j)
		j++
	}

	for k := range 5 {
		fmt.Println(k)
	}

	for {
		fmt.Println("loop")
		break
	}
}
