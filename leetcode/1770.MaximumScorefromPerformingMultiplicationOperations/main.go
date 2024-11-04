package main

import (
	"fmt"
	"math"
)

func dp(nums []int, multipliers []int, op, i int, memol, memor [][]int) int {
	if op == len(multipliers) {
		return 0
	}

	if memol[op][i] == math.MinInt {
		memol[op][i] = multipliers[op]*nums[i] + dp(nums, multipliers, op+1, i+1, memol, memor)
	}

	if memor[op][i] == math.MinInt {
		memor[op][i] = multipliers[op]*nums[len(nums)-1-(op-i)] + dp(nums, multipliers, op+1, i, memol, memor)
	}

	return max(
		memol[op][i],
		memor[op][i],
	)

}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func maximumScore(nums []int, multipliers []int) int {
	op := 0
	i := 0

	mult_len := len(multipliers)
	memol, memor := make([][]int, mult_len), make([][]int, mult_len)
	for i := range memol {
		sl := make([]int, mult_len)
		memol[i] = sl
		for j := range sl {
			sl[j] = math.MinInt
		}
	}

	for i := range memor {
		sl := make([]int, mult_len)
		memor[i] = sl
		for j := range sl {
			sl[j] = math.MinInt
		}
	}

	return dp(nums, multipliers, op, i, memol, memor)
}

func main() {
	// nums := []int{1, 2, 3}
	// multipliers := []int{3, 2, 1}
	// fmt.Println(maximumScore(nums, multipliers))
	nums := []int{-5, -3, -3, -2, 7, 1}
	multipliers := []int{-10, -5, 3, 4, 6}
	fmt.Println(maximumScore(nums, multipliers))

}
