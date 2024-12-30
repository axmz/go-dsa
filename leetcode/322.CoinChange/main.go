package main

import (
	"fmt"
	"math"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func coinChangeTD(coins []int, amount int) int {

	if len(coins) == 0 {
		return -1
	}

	memo := make(map[int]int)

	var dp func(a int) int
	dp = func(a int) int {
		if a < 0 {
			return -1
		}

		if a == 0 {
			return 0
		}

		if v, ok := memo[a]; ok {
			return v
		}

		minim := math.MaxInt

		for _, v := range coins {
			res := dp(a - v)
			if res >= 0 && res < minim {
				minim = min(minim, 1+res)
			}
		}

		if minim == math.MaxInt {
			minim = -1
		}

		memo[a] = minim
		return minim
	}

	return dp(amount)
}

func coinChange(coins []int, amount int) int {
	if len(coins) == 0 {
		return -1
	}

	if amount == 0 {
		return 0
	}

	res := math.MaxInt

	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt
	}

	for i := 1; i <= amount; i++ {
		res = math.MaxInt
		for c := range coins {
			prev := i - coins[c]
			if prev < 0 {
				dp[i] = min(dp[i], math.MaxInt)
			} else {
				if dp[prev] != math.MaxInt {
					dp[i] = min(dp[i], dp[prev]+1)
					res = min(res, dp[i])
				}
			}
		}
	}

	if res == math.MaxInt {
		return -1
	}
	return res
}

func main() {
	// coins := []int{2}
	// amount := 3
	// coins := []int{3}
	// amount := 2
	// coins := []int{1}
	// amount := 0
	// coins := []int{1, 2, 3}
	// amount := 6
	// coins := []int{1, 2, 5}
	// amount := 11
	// coins := []int{1, 2, 5, 10}
	// amount := 18
	// coins := []int{1, 2, 5}
	// amount := 8
	coins := []int{1, 25}
	amount := 2
	fmt.Println(coinChange(coins, amount))
}
