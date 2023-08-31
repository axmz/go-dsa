package main

import (
	"fmt"
)

func isVowel(b byte) bool {
	switch b {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	}
	return false
}

func countVowels(s []byte) int {
	t := 0
	for _, value := range s {
		if isVowel(value) {
			t++
		}
	}

	return t
}

func maxVowels(s string, k int) int {
	b := []byte(s)
	subs := b[0:k]
	vowelsCount := countVowels(subs)
	max := vowelsCount
	l, r := 1, k
	for r < len(s) {
		if isVowel(b[r]) {
			vowelsCount++
		}
		if isVowel(b[l-1]) {
			vowelsCount--
		}
		if vowelsCount > max {
			max = vowelsCount
		}
		if max == k {
			return max
		}
		l++
		r++
	}

	return max
}

func main() {
	// s := "abciiidef"
	// k := 3
	// s := "aeiou"
	// k := 2
	// s := "leetcode"
	// k := 3
	s := "weallloveyou"
	k := 7
	res := maxVowels(s, k)
	fmt.Println(res)
}
