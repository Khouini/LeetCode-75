package main

import (
	"fmt"
)

func uniqueOccurrences(arr []int) bool {
	occurences := make(map[int]int)
	for _, el := range arr {
		occurences[el]++
	}

	seen := make(map[int]bool)

	for _, el := range occurences {
		if seen[el] {
			return false
		}
		seen[el] = true
	}

	return true
}

func main() {
	fmt.Println(uniqueOccurrences([]int{1, 2, 2, 1, 1, 3}))
}
