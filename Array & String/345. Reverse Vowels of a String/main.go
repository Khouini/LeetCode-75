package main

import (
	"fmt"
)

func reverseVowels(s string) string {
	vowels := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		if isVowel(s[i]) {
			vowels = append(vowels, s[i])
		}
	}

	j := len(vowels) - 1
	result := []byte(s)
	for i := 0; i < len(s); i++ {
		if isVowel(s[i]) {
			result[i] = vowels[j]
			j--
		}
	}

	return string(result)
}

func isVowel(i byte) bool {
	if i >= 'A' && i <= 'Z' {
		i = i + 32
	}

	switch i {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	}

	return false
}

func main() {
	fmt.Println(reverseVowels("IceCreAm"))
}
