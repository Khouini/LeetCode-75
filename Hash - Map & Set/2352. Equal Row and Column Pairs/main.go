package main

import (
	"fmt"
	"strconv"
)

func equalPairs(grid [][]int) int {
	rows := make([]string, 0)
	cols := make(map[string]int)

	n := len(grid)

	for i := 0; i < n; i++ {
		row := ""
		col := ""
		for j := 0; j < n; j++ {
			if j > 0 {
				row += ","
				col += ","
			}
			row += strconv.Itoa(grid[i][j])
			col += strconv.Itoa(grid[j][i])
		}
		rows = append(rows, row)
		cols[col]++
	}

	count := 0
	for _, row := range rows {
		if _, exists := cols[row]; exists {
			count = count + cols[row]
		}
	}

	return count
}

func main() {
	fmt.Println(equalPairs([][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}}))                        // 1
	fmt.Println(equalPairs([][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}})) // 3
	fmt.Println(equalPairs([][]int{{3, 1, 2, 2}, {1, 4, 4, 4}, {2, 4, 2, 2}, {2, 5, 2, 2}})) //3
}
