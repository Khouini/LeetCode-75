package main

import "fmt"

func predictPartyVictory(senate string) string {
	R := make([]int, 0)
	D := make([]int, 0)
	killed := make(map[int]bool)
	for i := 0; i < len(senate); i++ {
		if senate[i] == 'R' {
			R = append(R, i)
		} else {
			D = append(D, i)
		}
	}

	skip := false
	i := 0
	j := 1
	for len(R) > 0 && len(D) > 0 {
		if killed[i] == true {
			skip = true
		}

		if skip == false {
			for senate[i] == senate[j] || killed[j] {
				j++
				if j == len(senate) {
					j = 0
				}
			}
			// thezare different now (we need t okill)
			killed[j] = true
			if senate[i] == 'R' {
				// KILL D
				D = D[1:]
			} else {
				// KILL R
				R = R[1:]
			}
		}

		i++
		if i == len(senate) {
			i = 0
		}
		j = i + 1
		if j == len(senate) {
			j = 0
		}
		skip = false
	}

	if len(R) == 0 && len(D) == 0 {
		panic("two zeros")
	}
	if len(R) == 0 {
		return "Dire"
	} else if len(D) == 0 {
		return "Radiant"
	} else {
		panic("no one is zero")
	}
}

func main() {
	tests := []struct {
		input    string
		expected string
	}{
		{"RD", "Radiant"},
		{"DR", "Dire"},
		{"RDD", "Dire"},
		{"DRR", "Radiant"},

		{"RRDD", "Radiant"},
		{"DDRR", "Dire"},
		{"RDRD", "Radiant"},
		{"DRDR", "Dire"},

		{"RRDDD", "Radiant"},
		{"DDRRR", "Dire"},
		{"RDRDD", "Dire"},
		{"DRDRR", "Radiant"},

		{"RRRDD", "Radiant"},
		{"DDDRR", "Dire"},
		{"RRRDDD", "Radiant"},
		{"DDDRRR", "Dire"},

		{"R", "Radiant"},
		{"D", "Dire"},
		{"RRRR", "Radiant"},
		{"DDDD", "Dire"},

		{"RDDR", "Radiant"},
		{"DRRD", "Dire"},
		{"RDDRR", "Radiant"},
		{"DRRDD", "Dire"},

		{"RRDRD", "Radiant"},
		{"DD RDR", "Dire"}, // remove this one if spaces are not allowed
	}

	for _, test := range tests {
		output := predictPartyVictory(test.input)

		fmt.Printf(
			"Input: %-8s | Output: %-8s | Expected: %-8s | Correct: %v\n",
			test.input,
			output,
			test.expected,
			output == test.expected,
		)
	}
}
