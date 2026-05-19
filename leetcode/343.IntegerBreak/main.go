package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func integerBreak(n int) int {
	dp := make([]int, n+1)

	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], max(j, dp[j])*max(i-j, dp[i-j]))
		}
	}
	return dp[n]
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
