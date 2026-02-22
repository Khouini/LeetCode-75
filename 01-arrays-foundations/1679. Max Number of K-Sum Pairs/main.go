package main

import (
	"fmt"
	"sort"
)

func maxOperations_BruteForce(nums []int, k int) int {
	n := len(nums)
	used := make([]bool, n)
	count := 0
	for i := 0; i < n; i++ {
		if used[i] {
			continue
		}
		for j := i + 1; j < n; j++ {
			if used[j] {
				continue
			}
			if nums[i]+nums[j] == k {
				used[i] = true
				used[j] = true
				count++
				break
			}
		}
	}
	return count
}

func maxOperations_TwoPointers(nums []int, k int) int {
	sort.Ints(nums)
	count := 0
	n := len(nums)
	L := 0
	R := n - 1

	for L < R {
		sum := nums[L] + nums[R]
		if sum == k {
			L++
			R--
			count++
		} else if sum > k {
			R--
		} else {
			L++
		}
	}
	return count
}

func maxOperations_HashMap(nums []int, k int) int {
	waitingRoom := make(map[int]int)
	count := 0

	for i := 0; i < len(nums); i++ {
		complement := k - nums[i]
		if waitingRoom[complement] > 0 {
			waitingRoom[complement]--
			count++
		} else {
			waitingRoom[nums[i]]++
		}
	}
	return count
}

func main() {
	fmt.Println(maxOperations_HashMap([]int{1, 2, 3, 4}, 5))    // 2
	fmt.Println(maxOperations_HashMap([]int{3, 1, 3, 4, 3}, 6)) // 1
}
