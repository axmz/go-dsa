package main

import (
	"fmt"
	"strings"
)

func removeStars(s string) string {
	stack := 0
	star := []byte("*")[0]
	b := []byte(s)

	for j := len(s) - 1; j >= 0; j-- {
		if s[j] == star {
			stack++
		} else if stack > 0 {
			b[j] = star
			stack--
		}
	}

	var sb strings.Builder
	for _, v := range b {
		if v != star {
			sb.WriteByte(v)
		}
	}

	return sb.String()
}

func main() {

	// t := "*"
	// t := "a"
	// t := "**"
	// t := "aa"
	// t := "*a"
	// t := "a*"
	// t := "**a"
	// t := "aa*"
	// t := "*a*"
	// t := "a*a"
	// t := "leet**cod*e"
	t := "erase*****"
	res := removeStars(t)
	fmt.Println(res)
}
