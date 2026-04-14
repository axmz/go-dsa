package main

import (
	"fmt"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minCost(costs [][]int) int {
	l := len(costs)
	dp := make([][3]int, l)
	dp[0][0], dp[0][1], dp[0][2] = costs[0][0], costs[0][1], costs[0][2]

	for i := 1; i < l; i++ {
		dp[i][0] = costs[i][0] + min(dp[i-1][1], dp[i-1][2])
		dp[i][1] = costs[i][1] + min(dp[i-1][0], dp[i-1][2])
		dp[i][2] = costs[i][2] + min(dp[i-1][0], dp[i-1][1])
	}

	last := l - 1
	return min(min(dp[last][0], dp[last][1]), dp[last][2])
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
