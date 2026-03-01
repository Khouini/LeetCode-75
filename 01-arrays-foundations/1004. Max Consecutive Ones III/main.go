package main

import "fmt"

// longestOnes returns the maximum number of consecutive 1s in the array
// if you can flip at most k zeros to ones.
func longestOnes(nums []int, k int) int {
	L := 0                   // Left pointer of the sliding window
	zerosCount := 0          // Number of zeros in the current window
	max := 0                 // Maximum length of valid window found
	n := len(nums)           // Length of the input array
	for R := 0; R < n; R++ { // R is the right pointer of the window
		if nums[R] == 0 {
			zerosCount++ // Increment zero count if current element is zero
		}
		// Shrink window from the left if zeros exceed k
		for zerosCount > k {
			if nums[L] == 0 {
				zerosCount-- // Decrement zero count when moving left pointer
			}
			L++ // Move left pointer forward
		}
		// Update max if the current window is larger
		if R-L+1 > max {
			max = R - L + 1
		}
	}

	return max // Return the largest window size found
}

func main() {
	// Example usage: should print the max consecutive 1s with at most 2 zero flips
	fmt.Println(longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2))
}
