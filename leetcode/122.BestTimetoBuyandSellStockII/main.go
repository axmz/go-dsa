package main

import "fmt"

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func maxProfit2(prices []int) int {
	held, sell, doNothing, prev := -prices[0], 0, 0, 0

	for i := 1; i < len(prices); i++ {
		sell = prices[i] + held
		doNothing = max(prev, held)
		held = max(held, -prices[i]+prev)
		prev = max(sell, doNothing)
	}

	return prev
}

func maxProfit(prices []int) int {
	maxProfit := 0
	i := 0

	for i < len(prices)-1 {
		valley, peak := 0, 0
		for i < len(prices)-1 && prices[i] >= prices[i+1] {
			i++
		}

		valley = prices[i]

		for i < len(prices)-1 && prices[i] <= prices[i+1] {
			i++
		}

		peak = prices[i]

		maxProfit += peak - valley
	}

	return maxProfit
}

func main() {
	// prices := []int{7, 1, 5, 3, 6, 4}
	prices := []int{1, 2, 3, 4, 5}
	// prices := []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(prices))
}
