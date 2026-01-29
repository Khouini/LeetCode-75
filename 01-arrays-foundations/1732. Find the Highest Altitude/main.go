package main

import (
	"fmt"
)

//	func largestAltitude(gain []int) int {
//		altitude := make([]int, len(gain)+1)
//		altitude[0] = 0
//		max := 0
//		for i := 1; i <= len(gain); i++ {
//			altitude[i] = altitude[i-1] + gain[i-1]
//			if altitude[i] > max {
//				max = altitude[i]
//			}
//		}
//		return max
//	}

func largestAltitude(gain []int) int {
	altitude := 0
	max := 0
	for i := 1; i <= len(gain); i++ {
		altitude = altitude + gain[i-1]
		if altitude > max {
			max = altitude
		}
	}
	return max
}

func main() {
	input := []int{-5, 1, 5, 0, -7}
	fmt.Println(largestAltitude(input))
}
