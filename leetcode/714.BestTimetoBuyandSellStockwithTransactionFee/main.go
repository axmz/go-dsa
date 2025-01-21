package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func maxProfit(prices []int, fee int) int {
	held, sell, doNothing, prev := -prices[0], 0, 0, 0

	for i := 1; i < len(prices); i++ {
		sell = prices[i] + held - fee
		doNothing = max(prev, held)
		held = max(held, -prices[i]+prev)
		prev = max(sell, doNothing)
	}

	return prev
}

func main() {
	// prices := []int{1, 3, 2, 8, 4, 9}
	// fee := 2
	prices := []int{1, 3, 7, 5, 10, 3}
	fee := 3
	fmt.Println(maxProfit(prices, fee))
}
