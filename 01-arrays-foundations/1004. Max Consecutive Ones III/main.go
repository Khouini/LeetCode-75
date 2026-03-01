package main

import "fmt"

func longestOnes(nums []int, k int) int {
	L := 0
	zerosCount := 0
	max := 0

	for R := 0; R < len(nums); R++ {
		if nums[R] == 0 {
			zerosCount++
		}

		// Shrink window until we have at most k zeros
		for zerosCount > k {
			if nums[L] == 0 {
				zerosCount--
			}
			L++
		}

		size := R - L + 1
		// Update best window size
		if size > max {
			max = size
		}
	}

	return max
}

func main() {
	fmt.Println(longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2))
}
