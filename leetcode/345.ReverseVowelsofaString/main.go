package main

import (
	"fmt"
)

func isVowel(ch byte) bool {
	switch ch {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	case 'a' - 32, 'e' - 32, 'i' - 32, 'o' - 32, 'u' - 32:
		return true
	}
	return false
}

func reverseVowels(s string) string {
	b := []byte(s)
	var isIVowel, isJVowel bool

	for i, j := 0, len(b)-1; i < j; {
		isIVowel = isVowel(b[i])
		isJVowel = isVowel(b[j])

		if isIVowel && isJVowel {
			b[i], b[j] = b[j], b[i]
			i++
			j--
			continue
		}

		if !isIVowel {
			i++
		}

		if !isJVowel {
			j--
		}
	}

	return string(b)
}

func main() {
	// s := "leetcode"
	s := "aA"
	res := reverseVowels(s)
	fmt.Println(res)
}
