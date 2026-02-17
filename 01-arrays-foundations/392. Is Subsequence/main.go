package main

func isSubsequence(s string, t string) bool {
	s_pointer := 0
	for t_pointer := 0; s_pointer < len(s) && t_pointer < len(t); t_pointer++ {
		if s[s_pointer] == t[t_pointer] {
			s_pointer++
		}
	}
	return s_pointer == len(s)
}

// func isSubsequence(s string, t string) bool {

// 	s_pointer := 0
// 	t_pointer := 0

// 	for s_pointer < len(s) && t_pointer < len(t) {
// 		if s[s_pointer] == t[t_pointer] {
// 			s_pointer++
// 		}
// 		t_pointer++
// 	}

// 	return s_pointer == len(s)

// }

func main() {
	s := "abc"
	t := "ahbgdc"
	println(isSubsequence(s, t)) // true

}
