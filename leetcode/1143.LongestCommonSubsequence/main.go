package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func longestCommonSubsequenceBottomUp1(text1 string, text2 string) int {
	dp := make([][]int, len(text2)+1)
	for i := range dp {
		dp[i] = make([]int, len(text1)+1)
	}

	for j := 1; j <= len(text2); j++ {
		for i := 1; i <= len(text1); i++ {
			if text1[i-1] == text2[j-1] {
				dp[j][i] = 1 + dp[j-1][i-1]
			} else {
				dp[j][i] = max(dp[j][i-1], dp[j-1][i])
			}
		}
	}

	return dp[len(text2)][len(text1)]
}

func longestCommonSubsequenceBottomUp2(text1 string, text2 string) int {
	prev := make([]int, len(text1)+1)
	cur := make([]int, len(text1)+1)

	for j := 1; j <= len(text2); j++ {
		for i := 1; i <= len(text1); i++ {
			if text1[i-1] == text2[j-1] {
				cur[i] = 1 + prev[i-1]
			} else {
				cur[i] = max(prev[i], cur[i-1])
			}
		}

		for k := range cur {
			prev[k] = cur[k]
			cur[k] = 0
		}
	}

	return prev[len(cur)-1]
}

func longestCommonSubsequenceTopDown(text1 string, text2 string) int {
	memo := make([][]int, len(text2))
	for i := range memo {
		memo[i] = make([]int, len(text1))
	}

	var lcs func(text1, text2 string, i, j int, memo [][]int) int
	lcs = func(text1, text2 string, i, j int, memo [][]int) int {
		if i == len(text1) || j == len(text2) {
			return 0
		}

		if memo[j][i] == -1 {
			return memo[j][i]
		}

		if text1[i] == text2[j] {
			memo[j][i] = 1 + lcs(text1, text2, i+1, j+1, memo)
		} else {
			memo[j][i] = max(lcs(text1, text2, i+1, j, memo), lcs(text1, text2, i, j+1, memo))
		}

		return memo[j][i]
	}

	return lcs(text1, text2, 0, 0, memo)
}

func main() {
	// text1 := "abcde"
	// text2 := "ace"
	text1 := "oxcpqrsvwf"
	text2 := "shmtulqrypy"
	fmt.Println(longestCommonSubsequenceTopDown(text1, text2))
	fmt.Println(longestCommonSubsequenceBottomUp1(text1, text2))
	fmt.Println(longestCommonSubsequenceBottomUp2(text1, text2))
}
