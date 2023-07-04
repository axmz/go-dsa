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

func recursion(cost []int, n int, memo map[int]int) int {

	if n < 0 {
		return 0
	}

	if _, ok := memo[n]; ok {
		return memo[n]
	}

	res := cost[n] + min(recursion(cost, n-1, memo), recursion(cost, n-2, memo))
	memo[n] = res

	return res
}

func minCostClimbingStairs(cost []int) int {

	memo := map[int]int{}
	return recursion(cost, len(cost)-1, memo)
	// if n < 0 return 0
	// return cost + min(cost len-1, cost len-2)

}

func minCostClimbingStairs2(cost []int) int {
	last2 := []int{cost[0], cost[1]}
	for i := 2; i < len(cost); i++ {
		current := cost[i] + min(last2[0], last2[1])
		last2[0] = last2[1]
		last2[1] = current
	}

	return min(last2[0], last2[1])
}

func main() {
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	res := minCostClimbingStairs(cost)
	res2 := minCostClimbingStairs2(cost)
	fmt.Println(res, res2)
}
