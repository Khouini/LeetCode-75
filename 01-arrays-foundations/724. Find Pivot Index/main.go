package main

import "log"

func pivotIndex(nums []int) int {
	prefix := make([]int, len(nums)+1)
	prefix[0] = 0

	for i := 1; i <= len(nums); i++ {
		prefix[i] = nums[i-1] + prefix[i-1]

	}

	suffix := make([]int, len(nums)+1)
	suffix[len(nums)] = 0

	for i := len(nums) - 1; i >= 0; i-- {
		suffix[i] = nums[i] + suffix[i+1]

	}

	for i := 0; i <= len(nums)-1; i++ {
		if prefix[i] == suffix[i+1] {
			return i
		}
	}
	return -1
}

func main() {
	log.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
	// log.Println(pivotIndex([]int{1, 2, 3}))
}
