package main

import "fmt"

func isMatch(s string, p string) bool {
	i, j := 0, 0
	star, match := -1, 0

	for i < len(s) {
		if j < len(p) && (p[j] == s[i] || p[j] == '?') {
			i++
			j++
		} else if j < len(p) && p[j] == '*' {
			star = j
			match = i
			j++
		} else if star != -1 {
			j = star + 1
			match++
			i = match
		} else {
			return false
		}
	}

	// remaining must all be '*'
	for j < len(p) && p[j] == '*' {
		j++
	}

	return j == len(p)
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
