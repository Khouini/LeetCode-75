package main

import (
	"fmt"
)

func longestSubarray(nums []int) int {
	n := len(nums)
	L := 0
	lastZeroIndex := -1

	max := 0
	for R := 0; R < n; R++ {

		if nums[R] == 0 {
			if lastZeroIndex != -1 {
				L = lastZeroIndex + 1
			}
			lastZeroIndex = R
		}

		size := R - L
		if size > max {
			max = size
		}

	}

	return max
}

func main() {
	fmt.Printf("got %d, expected %d\n", longestSubarray([]int{1, 1, 0, 1}), 3)
	fmt.Printf("got %d, expected %d\n", longestSubarray([]int{0, 1, 1, 1, 0, 1, 1, 0, 1}), 5)
	fmt.Printf("got %d, expected %d\n", longestSubarray([]int{1, 1, 1}), 2)
	fmt.Printf("got %d, expected %d\n", longestSubarray([]int{1, 0, 0, 0, 0}), 1)
	fmt.Printf("got %d, expected %d\n", longestSubarray([]int{1}), 0)
	fmt.Printf("got %d, expected %d\n", longestSubarray([]int{1, 1, 0, 0}), 2)
}
