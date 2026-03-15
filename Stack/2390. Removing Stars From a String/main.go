package main

import "fmt"

func removeStars(s string) string {
	var result []byte

	for i := 0; i < len(s); i++ {
		if s[i] == '*' {
			result = result[:len(result)-1]
		} else {
			result = append(result, s[i])
		}
	}
	return string(result)
}

func main() {
	fmt.Println(removeStars("leet**cod*e"))
	fmt.Println(removeStars("erase*****"))
}
