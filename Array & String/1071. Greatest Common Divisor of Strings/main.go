package main

import "fmt"

func gcdOfStrings(str1 string, str2 string) string {

	if str1+str2 != str2+str1 {
		return ""
	}

	lenX := gcd(len(str1), len(str2))

	return str1[0:lenX]
}

func gcd(a, b int) int {
	for b != 0 {
		r := a % b
		a = b
		b = r
	}
	return a
}

func main() {
	//fmt.Println(gcdOfStrings("ABABAB", "ABAB"))
	fmt.Println(gcdOfStrings("ABABAB", "ABAB"))
}
