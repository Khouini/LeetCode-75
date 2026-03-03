package main

import (
	"fmt"
)

func findDifference(nums1 []int, nums2 []int) [][]int {
	nums1Set := make(map[int]bool)
	nums2Set := make(map[int]bool)

	for _, num := range nums1 {
		nums1Set[num] = true
	}

	for _, num := range nums2 {
		nums2Set[num] = true
	}

	result1 := make([]int, 0)
	isPut1 := make(map[int]bool)
	for _, num := range nums1 {
		if nums2Set[num] == false && !isPut1[num] {
			result1 = append(result1, num)
			isPut1[num] = true
		}
	}

	result2 := make([]int, 0)
	isPut2 := make(map[int]bool)

	for _, num := range nums2 {
		if nums1Set[num] == false && !isPut2[num] {
			result2 = append(result2, num)
			isPut2[num] = true
		}
	}

	matrix := [][]int{result1, result2}

	return matrix
}

func main() {
	// fmt.Println(findDifference([]int{1, 2, 3}, []int{2, 4, 6}))
	fmt.Println(findDifference([]int{1, 2, 3, 3}, []int{1, 1, 2, 2}))

}
