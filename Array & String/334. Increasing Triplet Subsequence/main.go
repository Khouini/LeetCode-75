package main

import (
	"log"
	"math"
)

func increasingTriplet(nums []int) bool {
	leftArray := make([]int, len(nums))

	leftArray[0] = math.MaxInt
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < leftArray[i-1] {
			leftArray[i] = nums[i-1]
		} else {
			leftArray[i] = leftArray[i-1]
		}
	}

	rightArray := make([]int, len(nums))
	rightArray[len(nums)-1] = math.MinInt
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i+1] > rightArray[i+1] {
			rightArray[i] = nums[i+1]
		} else {
			rightArray[i] = rightArray[i+1]
		}
	}

	for i := 1; i < len(nums)-1; i++ {
		if leftArray[i] < nums[i] && nums[i] < rightArray[i] {
			return true
		}
	}
	return false
}

func main() {

	log.Println(increasingTriplet([]int{2, 1, 5, 0, 4, 6}))
}
