package main

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {
	curChar := chars[0]
	count := 1
	writeLen := 1
	groupStart := 0
	for i, ch := range chars {
		if i == 0 {
			continue
		}
		if ch != curChar {
			count = 1
			writeLen++
			curChar = ch
			groupStart = writeLen - 1
			chars[groupStart] = ch // write the character!
		} else {
			count++
			digits := strconv.Itoa(count)
			digitPos := groupStart + 1
			for _, c := range digits {
				chars[digitPos] = byte(c)
				digitPos++
			}
			writeLen = groupStart + 1 + len(digits)
		}
	}

	chars = chars[:writeLen]

	fmt.Println(string(chars))
	return writeLen
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
