package main

import "fmt"

func longestPalindrome(s string) string {
	expand := func(i, j int) string {
		l := i
		r := j
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
		}

		return s[l+1 : r]
	}

	ans := ""

	for i := 0; i < len(s); i++ {
		odd := expand(i, i)
		if len(odd) > len(ans) {
			ans = odd
		}
	}

	for i := 0; i < len(s); i++ {
		even := expand(i, i+1)
		if len(even) > len(ans) {
			ans = even
		}
	}

	return ans
}

func main() {
	s := "dxbaxbabd"
	fmt.Println(longestPalindrome(s))
}
