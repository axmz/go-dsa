package main

import "fmt"

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}

func minDistance(word1 string, word2 string) int {
	l1, l2 := len(word1)+1, len(word2)+1
	dp := make([][]int, l1)

	for i := range dp {
		dp[i] = make([]int, l2)
	}

	for i := range dp[0] {
		dp[0][i] = i
	}

	for i := range dp {
		dp[i][0] = i
	}

	for i := 1; i < l1; i++ {
		for j := 1; j < l2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(
					dp[i-1][j],
					dp[i][j-1],
					dp[i-1][j-1],
				) + 1
			}
		}
	}
	fmt.Println(dp)
	return dp[l1-1][l2-1]
}

func main() {
	word1, word2 := "ros", "horse"
	fmt.Println(minDistance(word1, word2))
}
