package main

import (
	"fmt"
)

func main() {
	// Make map string: int key values
	m1 := make(map[string]int)
	m1["key1"] = 1
	m1["key2"] = 2
	fmt.Println("m1: ", m1)

	// Retrieve values from map
	v1 := m1["key1"]
	v2 := m1["key2"]
	fmt.Println("Retrieved v1: ", v1)
	fmt.Println("Retrieved v2: ", v2)

	// Delete specific key from map
	delete(m1, "key1")
	fmt.Println("Removed key 1: ", m1)

	// Clear all key/values in map
	clear(m1)
	fmt.Println("Cleared m1: ", m1)

	// Determine if key is present in the map
	// Retrieving a non-existing key's value gives zero-val by default
	_, present := m1["key1"]
	fmt.Println("Present key? ", present)

	// Make a map in-line
	m2 := map[int]int{1: 2, 3: 4}
	fmt.Println("m2: ", m2)

	m3 := map[string]int{"hello": 3, "bye": 4, "exit": 0}
	fmt.Println("m3: ", m3)
}
