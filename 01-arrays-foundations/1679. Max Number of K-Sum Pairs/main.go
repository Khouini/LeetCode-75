package main

import "fmt"

func maxOperations(nums []int, k int) int {

	n := len(nums)
	i := 0
	j := 1
	ignoredIndexes := make([]int, 0, n)
	output := 0
	for i < n && j < n {
		isEnd := j == n-1

		iIsIgnored := false
		jIsIgnored := false
		for x := 0; x < len(ignoredIndexes); x++ {
			if ignoredIndexes[x] == i {
				iIsIgnored = true
			}
			if ignoredIndexes[x] == j {
				jIsIgnored = true
			}
		}

		if iIsIgnored {
			i++
			j = i + 1
			continue
		}

		if jIsIgnored {
			j++
			continue
		}

		isEqual := nums[i]+nums[j] == k
		if isEqual {
			output++
			ignoredIndexes = append(ignoredIndexes, i, j)
		}
		if isEnd || isEqual {
			i++
			j = i + 1
			continue
		}

		j++

	}
	return output
}

func main() {
	// fmt.Println(maxOperations([]int{1, 2, 3, 4}, 5))
	fmt.Println(maxOperations([]int{3, 1, 3, 4, 3}, 6))
}
