package main

import (
	"fmt"
	"sort"
)

func maxOperations(nums []int, k int) int {
	sort.Ints(nums)
	L, R := 0, len(nums)-1
	count := 0

	for L < R {
		sum := nums[L] + nums[R]
		if sum == k {
			count++
			L++
			R--
		} else if sum < k {
			L++
		} else {
			R--
		}
	}

	return count
}

func main() {
	fmt.Println(maxOperations([]int{1, 2, 3, 4}, 5))    // 2
	fmt.Println(maxOperations([]int{3, 1, 3, 4, 3}, 6)) // 1
}
