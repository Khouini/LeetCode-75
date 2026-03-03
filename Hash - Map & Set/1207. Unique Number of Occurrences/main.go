package main

import (
	"fmt"
)

func uniqueOccurrences(arr []int) bool {
	occurences := make(map[int]int)
	for _, el := range arr {
		occurences[el]++
	}

	seen := make(map[int]struct{})

	for _, el := range occurences {
		if _, exists := seen[el]; exists {
			return false
		}
		seen[el] = struct{}{}
	}

	return true
}

func main() {
	fmt.Println(uniqueOccurrences([]int{1, 2, 2, 1, 1, 3}))
}
