package main

import "fmt"

func numWays(n int, k int) int {
	if n == 1 {
		return k
	}
	if n == 2 {
		return k * k
	}
	dp := make([]int, n)
	dp[0] = k
	dp[1] = k * k
	for i := 2; i < n; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) * (k - 1)
	}
	return dp[n-1]
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
