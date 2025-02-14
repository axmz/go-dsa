package main

import "fmt"

func maxProfit(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		profit := prices[i] - minPrice
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
		if profit > maxProfit {
			maxProfit = profit
		}
	}

	return maxProfit
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}
