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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func mincostTickets(days []int, costs []int) int {
	// Create a set of travel days for O(1) lookups
	daySet := make(map[int]bool)
	for _, day := range days {
		daySet[day] = true
	}

	lastDay := days[len(days)-1]
	// dp is one dimension not three
	dp := make([]int, lastDay+1)

	for i := 1; i <= lastDay; i++ {
		if !daySet[i] {
			// If it's not a travel day, cost is the same as the previous day
			dp[i] = dp[i-1]
		} else {
			// Calculate cost for 1-day, 7-day, and 30-day passes
			// Nice technique, instead of calculating forward you calculate up to i
			// Also nice way to handle out-of-bounds by using max(0, i-x)
			cost1 := dp[max(0, i-1)] + costs[0]
			cost7 := dp[max(0, i-7)] + costs[1]
			cost30 := dp[max(0, i-30)] + costs[2]

			// Take the minimum of the three options
			dp[i] = min(cost1, min(cost7, cost30))
		}
	}

	return dp[lastDay]
}

func mincostTicketsTopDown(days []int, costs []int) int {
	memo := make(map[int]int)
	end := days[len(days)-1]

	daySet := make(map[int]bool)
	for _, day := range days {
		daySet[day] = true
	}

	var dfs func(day int) int
	dfs = func(day int) int {
		if day > end {
			return 0
		}

		if v, ok := memo[day]; ok {
			return v
		}

		// nice technique to skip non-travel days
		if !daySet[day] {
			memo[day] = dfs(day + 1)
			return memo[day]
		}

		cost1 := costs[0] + dfs(day+1)
		cost7 := costs[1] + dfs(day+7)
		cost30 := costs[2] + dfs(day+30)

		memo[day] = min(cost1, min(cost7, cost30))
		return memo[day]
	}

	return dfs(days[0])
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
