package main

import "fmt"

func isMatch2(s string, p string) bool {
	ls, lp := len(s), len(p)

	dp := make([][]bool, ls+1)
	for i := range dp {
		dp[i] = make([]bool, lp+1)
	}

	dp[0][0] = true

	for j := 2; j <= lp; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}

	for i := 1; i <= ls; i++ {
		for j := 1; j <= lp; j++ {
			if p[j-1] == '*' {

				dp[i][j] = dp[i][j-2]

				if s[i-1] == p[j-2] || p[j-2] == '.' {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			} else if s[i-1] == p[j-1] || p[j-1] == '.' {

				dp[i][j] = dp[i-1][j-1]
			}
		}
	}

	return dp[ls][lp]
}

func isMatch(s string, p string) bool {
	memo := make(map[[2]int]bool)
	var dfs func(i, j int) bool
	dfs = func(i, j int) bool {
		if j == len(p) {
			return i == len(s)
		}

		if v, ok := memo[[2]int{i, j}]; ok {
			return v
		}

		match := i < len(s) && (s[i] == p[j] || p[j] == '.')
		if j+1 < len(p) && p[j+1] == '*' {
			memo[[2]int{i, j}] = dfs(i, j+2) || (match && dfs(i+1, j))
			return memo[[2]int{i, j}]
		}

		if match {
			memo[[2]int{i, j}] = dfs(i+1, j+1)
			return memo[[2]int{i, j}]
		}

		memo[[2]int{i, j}] = false
		return false
	}

	return dfs(0, 0)
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
