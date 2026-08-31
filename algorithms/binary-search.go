package main

import "fmt"

func binarySearch(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	minIdx := 0
	maxIdx := len(nums) - 1

	for minIdx <= maxIdx {
		midIdx := minIdx + (maxIdx-minIdx)/2

		if nums[midIdx] == target {
			return midIdx
		} else if nums[midIdx] < target {
			minIdx = midIdx + 1
		} else {
			maxIdx = midIdx - 1
		}

	}

	return -1
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	target := 10

	elementPosition := binarySearch(nums, target)
	if elementPosition == -1 {
		fmt.Printf("Value (%d) was not present in the array.\n", target)
	} else {
		fmt.Printf("Value (%d) was found at position (%d) in the array.\n", target, elementPosition)
	}
}
