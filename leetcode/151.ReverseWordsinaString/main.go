package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	s = " " + s // this makes it cleaner
	res := make([]byte, 0, len(s))

	for l, r := len(s)-1, len(s)-1; l >= 0; {

		if s[l] == 32 {
			l--
			r = l
			continue
		}

		for s[l] != 32 && l > 0 {
			l--
		}

		// This solution is very ugly
		// if l == 0 && s[l] != 32 {
		// 	res = append(res, s[l:r+1]...)
		// 	break
		// } else {
		res = append(res, s[l+1:r+1]...)
		res = append(res, ' ')
		// }
	}

	if res[len(res)-1] == 32 {
		return string(res[:len(res)-1])
	}
	return string(res)
}

func reverseWordsBuiltIn(s string) string {
	words := strings.Fields(s)

	start := 0
	end := len(words) - 1

	for start < end {
		words[start], words[end] = words[end], words[start]
		start++
		end--
	}

	return strings.Join(words, " ")
}

func main() {
	s := " a good   example"
	// s := "the sky is blue"
	// s := "  hello world  "
	res := reverseWords(s)
	fmt.Print(res)
	fmt.Printf("|")
}
