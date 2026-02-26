package main

import (
	"fmt"
	"math"
)

func maxVowels(s string, k int) int {
	maxSum := math.MinInt
	sum := 0
	for i := 0; i < len(s); i++ {
		sum = sum + vowelValue(s[i])

		if i >= k {
			sum = sum - vowelValue(s[i-k])
		}
		if i >= k-1 {
			if sum > maxSum {
				maxSum = sum
			}
		}
	}
	return maxSum
}

func vowelValue(c byte) int {
	switch c {
	case 'a', 'e', 'i', 'o', 'u':
		return 1
	}
	return 0
}

func main() {
	fmt.Println(maxVowels("leetcode", 3))
}
