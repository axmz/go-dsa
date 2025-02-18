package main

import "fmt"

func compress(s string) string {
	l := len(s)

	if l == 1 {
		return "1" + s
	}

	count := 1
	res := ""
	for i := 1; i < l; i++ {
		if s[i-1] != s[i] {
			res += fmt.Sprintf("%d", count) + string(s[i-1])
			count = 1
		} else {
			count++
		}
	}
	res += fmt.Sprintf("%d", count) + string(s[l-1])

	return res
}

func countAndSay(n int) string {
	s := "1"

	for i := 1; i < n; i++ {
		s = compress(s)
	}

	return s
}

func main() {
	x := 4
	fmt.Println(compress("2223"))
	fmt.Println(countAndSay(x))
}
