package main

import "fmt"

func isInterleave2(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	dp := make([][]bool, len(s1)+1)
	for i := range dp {
		dp[i] = make([]bool, len(s2)+1)
	}

	dp[0][0] = true

	for i := 1; i <= len(s1); i++ {
		dp[i][0] = dp[i-1][0] && s1[i-1] == s3[i-1]
	}

	for j := 1; j <= len(s2); j++ {
		dp[0][j] = dp[0][j-1] && s2[j-1] == s3[j-1]
	}

	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			k := i + j
			dp[i][j] = (dp[i-1][j] && s1[i-1] == s3[k-1]) || (dp[i][j-1] && s2[j-1] == s3[k-1])
		}
	}

	return dp[len(s1)][len(s2)]
}

func isInterleave3(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	memo := make(map[[2]int]bool)
	var dfs func(i, j int) bool
	dfs = func(i, j int) bool {
		if i == len(s1) && j == len(s2) {
			return true
		}

		if v, ok := memo[[2]int{i, j}]; ok {
			return v
		}

		k := i + j
		if i < len(s1) && s1[i] == s3[k] && dfs(i+1, j) {
			memo[[2]int{i, j}] = true
			return true
		}

		if j < len(s2) && s2[j] == s3[k] && dfs(i, j+1) {
			memo[[2]int{i, j}] = true
			return true
		}

		memo[[2]int{i, j}] = false
		return false
	}

	return dfs(0, 0)
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	n := len(s1)
	m := len(s2)
	if n+m != len(s3) {
		return false
	}

	dp := make([]bool, m+1)
	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			k := i + j
			if i == 0 && j == 0 {
				dp[j] = true
			} else if i == 0 {
				dp[j] = dp[j-1] && s2[j-1] == s3[k-1]
			} else if j == 0 {
				dp[j] = dp[j] && s1[i-1] == s3[k-1]
			} else {
				dp[j] = (dp[j] && s1[i-1] == s3[k-1]) || (dp[j-1] && s2[j-1] == s3[k-1])
			}
		}
	}
	return dp[m]
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
