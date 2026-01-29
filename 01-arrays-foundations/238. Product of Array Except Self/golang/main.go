package main

import "fmt"

func productExceptSelf_BruteForce(nums []int) []int {
	result := make([]int, len(nums))
	product := 1
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if j == i {
				continue
			}
			product = product * nums[j]
		}

		result[i] = product
		product = 1
	}

	return result
}

func productExceptSelf(nums []int) []int {
	n := len(nums)

	prefix := make([]int, n)
	prefix[0] = 1
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] * nums[i-1]
	}

	suffix := make([]int, n)
	suffix[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		suffix[i] = suffix[i+1] * nums[i+1]
	}

	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = prefix[i] * suffix[i]
	}

	return result
}

func main() {
	// This is a placeholder for the main function.
	// result := productExceptSelf_BruteForce([]int{1, 2, 3, 4})
	result := productExceptSelf([]int{1, 2, 3, 4})
	fmt.Println(result)
}
