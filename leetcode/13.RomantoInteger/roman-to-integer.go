package main

import "fmt"

func romanToInt(s string) int {
	table := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	l := len(s)
	res := table[string(s[l-1])]
	for j, i := 0, 1; i < l; j, i = j+1, i+1 {
		I := table[string(s[i])]
		J := table[string(s[j])]
		println(J, I, res)
		if J < I {
			res -= J
		} else {
			res += J
		}
	}

	return res
}

func main() {
	fmt.Println("Roman to Integer:", romanToInt("III"))
}
