package main

import (
	"fmt"
)

func isSubsequence(s string, t string) bool {
	if s == "" {
		return true
	}
	for i, j := 0, 0; i < len(t) && j < len(s); i++ {
		if t[i] == s[j] {
			j++
		}
		if j == len(s) {
			return true
		}
	}

	return false
}

func main() {
	s := ""
	t := "ahbgdx"

	res := isSubsequence(s, t)
	fmt.Println(res)
}
