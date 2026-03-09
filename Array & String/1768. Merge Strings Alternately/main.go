package main

import "fmt"

func mergeAlternately(word1 string, word2 string) string {
	minLength := 0
	var bigWord string
	if len(word1) > len(word2) {
		minLength = len(word2)
		bigWord = word1
	} else {
		minLength = len(word1)
		bigWord = word2
	}
	result := make([]byte, 0)
	for i := 0; i < minLength; i++ {
		result = append(result, word1[i])
		result = append(result, word2[i])
	}

	for i := minLength; i < len(bigWord); i++ {
		result = append(result, bigWord[i])
	}

	return string(result)
}

func main() {
	fmt.Println(mergeAlternately("abc", "pqr"))
	fmt.Println(mergeAlternately("ab", "pqrs"))
	fmt.Println(mergeAlternately("abcd", "pq"))
}
