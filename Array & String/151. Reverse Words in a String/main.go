package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	word := make([]byte, 0)
	words := make([]string, 0)

	for i := 0; i <= len(s); i++ {
		if i == len(s) || s[i] == ' ' {
			if len(word) > 0 {
				words = append(words, string(word))
			}
			word = word[:0]
		} else {
			word = append(word, s[i])
		}
	}

	var result strings.Builder
	for i := len(words) - 1; i >= 0; i-- {
		if i < len(words)-1 {
			result.WriteByte(' ')
		}
		result.WriteString(words[i])
	}

	return result.String()
}

func main() {

	fmt.Println(reverseWords("  hello world  "))
	fmt.Println(reverseWords("a good   example"))
}
