package main

import "fmt"

func main() {
	// Iterate over the elements of a slice/array
	nums := []int{1,2,3,4,5}
	for _, num := range nums {
		fmt.Println(num)
	}


	// Range provides (index, value) for each entry
	for i, num := range(nums) {
		fmt.Printf("index: %d, num: %d\n", i, num)
	}

	// Range over map iterates over key/value pairs
	map1 := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range map1 {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// Iterate over only the keys of a map
	for k := range map1 {
		fmt.Println("key:", k)
	}

	 // Range on Strings iterates over Unicode code points
	// 1st value is the starting byte index of the rune, 2nd the rune itself
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
