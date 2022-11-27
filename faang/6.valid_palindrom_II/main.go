package main

import (
	"fmt"
)

func simplePalindrome(s string) bool {
	var l, r = 0, len(s) - 1

	for l < r {

		if s[l] == s[r] {
			l++
			r--

		} else {
			return false
		}
	}

	return true

}

func validPalindrome(s string) bool {
	if len(s) <= 2 {
		return true
	}

	var l, r = 0, len(s) - 1

	for l < r {

		if s[l] == s[r] {
			l++
			r--

		} else {
			if simplePalindrome(s[l:r]) || simplePalindrome(s[l+1:r+1]) {
				return true
			}

			return false
		}
	}

	return true
}

func main() {
	res := validPalindrome("abc")
	// res := validPalindrome("abccdba")
	// res := simplePalindrome("ababba")
	fmt.Println("Result: ", res)
}
