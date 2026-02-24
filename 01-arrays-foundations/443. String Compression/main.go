package main

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {
	lastCharacter := chars[0]
	count := 1
	result := 1
	characterGroupIndex := 0
	for i, el := range chars {
		// fmt.Printf("%c\n", el)
		if i == 0 {
			continue
		}
		if el != lastCharacter {
			count = 1
			result++
			lastCharacter = el
			characterGroupIndex = result - 1
			chars[characterGroupIndex] = el // write the character!
		} else {
			count++
			isCount2 := count == 2
			isModulo10 := count%10 == 0
			if isCount2 || isModulo10 {
				result++
			}
			writeIndex := characterGroupIndex + 1
			for _, c := range strconv.Itoa(count) {
				chars[writeIndex] = byte(c)
				writeIndex++ // advance after each digit
			}
		}

	}

	chars = chars[:result]

	fmt.Println(string(chars))
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
