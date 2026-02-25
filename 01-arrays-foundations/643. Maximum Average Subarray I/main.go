package main

import (
	"fmt"
	"math"
)

// func findMaxAverage(nums []int, k int) float64 {
// 	maxAverage := -math.MaxFloat64
// 	for i := 0; i < len(nums); i++ {
// 		isExceded := i+k > len(nums)
// 		if isExceded {
// 			break
// 		}
// 		limit := i + k
// 		sum := float64(0)
// 		for j := i; j < limit; j++ {
// 			sum = sum + float64(nums[j])
// 		}
// 		average := sum / float64(k)
// 		if average > maxAverage {
// 			maxAverage = average
// 		}
// 	}

// 	return maxAverage
// }

func findMaxAverage(nums []int, k int) float64 {
	currentSum := float64(0)
	maxAverage := -math.MaxFloat64
	for i := 0; i < len(nums); i++ {
		currentSum = currentSum + float64(nums[i])
		if i >= k {
			currentSum = currentSum - float64(nums[i-k])
		}
		if i >= k-1 {
			average := currentSum / float64(k)
			if average > maxAverage {
				maxAverage = average
			}
		}
	}

	return maxAverage
}

func main() {
	// fmt.Println(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4))
	fmt.Println(findMaxAverage([]int{5}, 1))
}
