package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func maxProfit2(prices []int) int {

	memo := make([][2][2]int, len(prices))
	for i := range memo {
		memo[i] = [2][2]int{{-1, -1}, {-1, -1}}
	}

	var dp func(i, cooldown, stock int) int
	dp = func(i, cooldown, stock int) int {
		if i == len(prices) {
			return 0
		}

		if val := memo[i][cooldown][stock]; val != -1 {
			return val
		}

		var doSomething, doNothing int

		if stock == 0 && cooldown == 0 {
			doSomething = -prices[i] + dp(i+1, 0, 1)
		} else if stock == 1 {
			doSomething = prices[i] + dp(i+1, 1, 0)
		}

		doNothing = dp(i+1, 0, stock)

		memo[i][cooldown][stock] = max(doSomething, doNothing)
		return memo[i][cooldown][stock]
	}

	return dp(0, 0, 0)
}

func maxProfit(prices []int) int {

	dp := make([][2][2]int, len(prices)+1)

	for i := len(prices) - 1; i >= 0; i-- {
		for cooldown := 0; cooldown < 2; cooldown++ {
			for stock := 0; stock < 2; stock++ {
				var doNothing, doSomething int

				if stock == 0 && cooldown == 0 {
					doSomething = -prices[i] + dp[i+1][0][1]
				} else if stock == 1 {
					doSomething = prices[i] + dp[i+1][1][0]
				}

				doNothing = dp[i+1][0][stock]

				dp[i][cooldown][stock] = max(max(doSomething, doNothing), dp[i][cooldown][stock])
			}
		}
	}

	return dp[0][0][0]
}

func main() {
	prices := []int{1, 2, 3, 0, 2}
	fmt.Println(maxProfit(prices))
}
