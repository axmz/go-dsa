package main

import "fmt"

func reverseWords(s string) string {
	b := []byte(s)

	for i, j := 0, 0; j < len(s); j++ {
		if b[j] == ' ' {
			continue
		}

		if j+1 == len(s) || b[j+1] == ' ' {
			for l, r := i, j; l < r; l, r = l+1, r-1 {
				b[l], b[r] = b[r], b[l]
			}
			i = j + 2
		}
	}

	return string(b)
}

func main() {
	s := "Let's take LeetCode contest"
	fmt.Println(reverseWords(s))
}
