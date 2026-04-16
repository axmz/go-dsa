package main

import "fmt"

func numRollsToTarget2(n int, k int, target int) int {
	mod := 1_000_000_007
	memo := make([][]int, n+1)
	for i := range memo {
		memo[i] = make([]int, target+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var backtrack func(diceLeft, targetLeft int) int
	backtrack = func(diceLeft, targetLeft int) int {
		if diceLeft == 0 && targetLeft == 0 {
			return 1
		}

		if targetLeft < 0 || diceLeft < 0 {
			return 0
		}

		if v := memo[diceLeft][targetLeft]; v != -1 {
			return v
		}

		count := 0
		for i := 1; i <= k; i++ {
			count += backtrack(diceLeft-1, targetLeft-i)
			count %= mod
		}

		memo[diceLeft][targetLeft] = count
		return count
	}

	return backtrack(n, target)
}

func numRollsToTarget3(n int, k int, target int) int {
	mod := 1_000_000_007
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, target+1)
	}

	dp[0][0] = 1
	for r := 1; r <= n; r++ {
		for t := 1; t <= target; t++ {
			for face := 1; face <= k; face++ {
				if t-face >= 0 {
					dp[r][t] += dp[r-1][t-face]
					dp[r][t] %= mod
				}
			}
		}
	}

	return dp[n][target]
}

func numRollsToTarget(n int, k int, target int) int {
	mod := 1_000_000_007
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, target+1)
	}

	dp[0][0] = 1
	for r := 1; r <= n; r++ {
		for t := target; t >= 1; t-- {
			for face := 1; face <= k; face++ {
				if t-face >= 0 {
					dp[r][t] += dp[r-1][t-face]
					dp[r][t] %= mod
				}
			}
		}
	}

	return dp[n][target]
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
