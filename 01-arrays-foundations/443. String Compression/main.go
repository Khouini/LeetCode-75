package main

import "fmt"

func compress(chars []byte) int {
	lastCharacter := chars[0]
	count := 1
	result := 1
	for i, el := range chars {
		// fmt.Printf("%c\n", el)
		if i == 0 {
			continue
		}
		if el != lastCharacter {
			count = 1
			result++
			lastCharacter = el
		} else {
			count++
			isCount2 := count == 2
			isModulo10 := count%10 == 0
			if isCount2 || isModulo10 {
				result++
			}
		}

	}

	return result
}

func main() {
	s1 := "aabbccc"
	fmt.Println(compress([]byte(s1)))
	s2 := "a"
	fmt.Println(compress([]byte(s2)))
	s3 := "abbbbbbbbbbbb"
	fmt.Println(compress([]byte(s3)))
	s4 := "aaaaaaaaaa"
	fmt.Println(compress([]byte(s4)))
}
