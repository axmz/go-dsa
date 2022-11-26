package main

import "fmt"

func LengthOfLongestSubstring(s string) int {
	var length, L, R int = 0, 0, 0

	seenChars := make(map[rune]int)

	updateR := func(ch rune, i int) {
		R++
		seenChars[ch] = i
	}

	updateL := func(newL int, ch rune, i int) {
		L = newL
		seenChars[ch] = i
	}

	for i, ch := range s {
		if seenAtIndex, wasSeen := seenChars[ch]; wasSeen {
			if seenAtIndex < L {
				updateR(ch, i)
			} else {
				updateL(seenAtIndex+1, ch, i)
			}
		} else {
			updateR(ch, i)
		}

		if i-L >= length {
			length++
		}
	}

	return length
}

func main() {
	// res := LengthOfLongestSubstring("abcbdaac")
	// res := LengthOfLongestSubstring("pwwkew")
	res := LengthOfLongestSubstring("bbtablud")
	fmt.Println("Result: ", res)
}
