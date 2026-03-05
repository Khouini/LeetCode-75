package main

import (
	"fmt"
	"sort"
)

func closeStrings(word1 string, word2 string) bool {
	word1Frequencies := make(map[byte]int)
	word2Frequencies := make(map[byte]int)

	for i := 0; i < len(word1); i++ {
		word1Frequencies[word1[i]]++
	}

	for i := 0; i < len(word2); i++ {
		word2Frequencies[word2[i]]++
	}

	if len(word1Frequencies) != len(word2Frequencies) {
		return false
	}

	w1Counts := make([]int, 0)
	w2Counts := make([]int, 0)
	for w1 := range word1Frequencies {
		if _, exists := word2Frequencies[w1]; !exists {
			return false
		}
		w1Counts = append(w1Counts, word1Frequencies[w1])
		w2Counts = append(w2Counts, word2Frequencies[w1])
	}

	sort.Ints(w1Counts)
	sort.Ints(w2Counts)

	for i := 0; i < len(w1Counts); i++ {
		if w1Counts[i] != w2Counts[i] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(closeStrings("cabbba", "abbccc"))
	fmt.Println(closeStrings("a", "aa"))
}
