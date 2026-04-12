package main

import "fmt"

func numDistinct2(s string, t string) int {
	memo := make(map[[2]int]int)
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if j == len(t) {
			return 1
		}
		if i == len(s) {
			return 0
		}

		if v, ok := memo[[2]int{i, j}]; ok {
			return v
		}

		count := dfs(i+1, j)
		if s[i] == t[j] {
			count += dfs(i+1, j+1)
		}

		memo[[2]int{i, j}] = count
		return count
	}

	return dfs(0, 0)
}

func numDistinct3(s string, t string) int {
	dp := make([][]int, len(s)+1)
	for i := range dp {
		dp[i] = make([]int, len(t)+1)
	}

	for i := 0; i <= len(s); i++ {
		dp[i][0] = 1
	}

	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(t); j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[len(s)][len(t)]
}

func numDistinct(s string, t string) int {
	m, n := len(s), len(t)
	if m < n {
		return 0
	}
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= m; i++ {
		for j := min(i, n); j >= 1; j-- {
			if s[i-1] == t[j-1] {
				dp[j] += dp[j-1]
			}
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
