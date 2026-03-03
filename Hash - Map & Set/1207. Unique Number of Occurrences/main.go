package main

import (
	"fmt"
	"sort"
)

func uniqueOccurrences(arr []int) bool {
	occurences := make(map[int]int)
	values := []int{}
	n := len(arr)
	for i := 0; i < n; i++ {
		occurences[arr[i]]++
	}
	for _, v := range occurences {
		values = append(values, v)
	}

	sort.Ints(values)

	for i := 0; i < len(values)-1; i++ {
		if values[i] == values[i+1] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(uniqueOccurrences([]int{1, 2, 2, 1, 1, 3}))
}
