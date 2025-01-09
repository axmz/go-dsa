package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func maxProfit(k int, prices []int) int {

	memo := make(map[string]int)

	var dp func(i, kb, ks, stock int) int
	dp = func(i, kb, ks, stock int) int {
		if i == len(prices) || ks == 0 {
			return 0
		}

		key := fmt.Sprintf("%d-%d-%d-%d", i, kb, ks, stock)
		if val, exists := memo[key]; exists {
			return val
		}

		var dp1, dp2, dp3 int

		if kb > 0 && stock == 0 {
			dp1 = dp(i+1, kb-1, ks, 1) - prices[i]
		}

		dp2 = dp(i+1, kb, ks, stock)

		if ks > 0 && stock == 1 {
			dp3 = dp(i+1, kb, ks-1, 0) + prices[i]
		}

		memo[key] = max(max(dp1, dp2), dp3)
		return memo[key]
	}

	return dp(0, k, k, 0)
}

func main() {
	// k := 2
	// prices := []int{2, 4, 1}
	// k := 2
	// prices := []int{3, 2, 6, 5, 0, 3}
	k := 7
	prices := []int{48, 12, 60, 93, 97, 42, 25, 64, 17, 56, 85, 93, 9, 48, 52, 42, 58, 85, 81, 84, 69, 36, 1, 54, 23, 15, 72, 15, 11, 94}
	fmt.Println(maxProfit(k, prices))
}
