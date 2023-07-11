package main

import (
	"fmt"
	"strings"
)

func max(a, b int) int {

	if a > b {
		return a
	}
	return b
}

func mergeAlternately(word1 string, word2 string) string {
	var sb strings.Builder
	l1, l2 := len(word1), len(word2)
	m := max(l1, l2)
	for i := 0; i < m; i++ {
		if i < l1 {
			sb.WriteByte(word1[i])
		}
		if i < l2 {
			sb.WriteByte(word2[i])
		}
	}

	return sb.String()
}

// There is also double pointers approach
// This approach doesn't work for utf8
func main() {
	word1 := "abc22"
	word2 := "pqr😒s"
	res := mergeAlternately(word1, word2)
	fmt.Println(res)
}
