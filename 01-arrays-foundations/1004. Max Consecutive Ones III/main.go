package main

import "fmt"

// longestOnes returns the maximum number of consecutive 1s in the array
// if you can flip at most k zeros to ones.
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

		// Update best window size
		if R-L+1 > max {
			max = R - L + 1
		}
	}

	return max
}

func main() {
	// Example usage: should print the max consecutive 1s with at most 2 zero flips
	fmt.Println(longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2))
}
