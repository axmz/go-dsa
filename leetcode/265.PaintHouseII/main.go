package main

import (
	"fmt"
	"math"
)

func minSlice(nums []int) int {
	m := math.MaxInt
	for _, n := range nums {
		if n < m {
			m = n
		}
	}
	return m
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minCostII(costs [][]int) int {
	k := len(costs[0])
	dp := make([][]int, len(costs))
	for i := range dp {
		dp[i] = make([]int, k)
	}
	copy(dp[0], costs[0])
	for i := 1; i < len(costs); i++ {
		for c := 0; c < len(costs[i]); c++ {
			cost := costs[i][c]
			dp[i][c] = cost + min(minSlice(dp[i-1][:c]), minSlice(dp[i-1][c+1:]))
		}
	}
	last := len(dp) - 1
	return minSlice(dp[last])
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
