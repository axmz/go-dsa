package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	var alphanumeric strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			alphanumeric.WriteRune(r)
		}
	}

	s = alphanumeric.String()
	s = strings.ToLower(s)
	s = strings.TrimSpace(s)

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}

	return true
}

func main() {
	s := "baaab"
	fmt.Println(isPalindrome(s))
}
