package main

import "fmt"

func maxAreaBruteForce(height []int) int {
	maxWater := 0
	n := len(height)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			width := j - i
			h := min(height[i], height[j])
			water := width * h
			maxWater = max(maxWater, water)
		}
	}
	return maxWater
}

func maxArea(height []int) int {
	n := len(height)
	L := 0
	R := n - 1
	maxArea := 0
	for R > L {
		h := min(height[L], height[R])
		width := R - L
		area := width * h
		maxArea = max(maxArea, area)

		if height[L] > height[R] {
			R = R - 1
		} else {
			L = L + 1
		}
	}

	return maxArea
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
