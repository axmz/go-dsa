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

func minCost(houses []int, cost [][]int, m int, n int, target int) int {
	memo := make(map[[3]int]int)

	var paint func(houseIdx, prevColor, neighborhoods int) int
	paint = func(houseIdx, prevColor, neighborhoods int) int {
		if neighborhoods > target {
			return math.MaxInt32
		}

		key := [3]int{houseIdx, prevColor, neighborhoods}
		if val, exists := memo[key]; exists {
			return val
		}

		if houseIdx == m {
			if neighborhoods == target {
				return 0
			}
			return math.MaxInt32
		}

		if houses[houseIdx] != 0 {
			newNeighborhoods := neighborhoods
			if houses[houseIdx] != prevColor {
				newNeighborhoods++
			}
			return paint(houseIdx+1, houses[houseIdx], newNeighborhoods)
		}

		minCost := math.MaxInt32

		// nice technqieu to start the count from 1
		for color := 1; color <= n; color++ {
			newNeighborhoods := neighborhoods
			if color != prevColor {
				newNeighborhoods++
			}
			costToPaint := cost[houseIdx][color-1]
			totalCost := costToPaint + paint(houseIdx+1, color, newNeighborhoods)
			minCost = min(minCost, totalCost)
		}

		memo[key] = minCost
		return minCost
	}

	result := paint(0, 0, 0)
	if result == math.MaxInt32 {
		return -1
	}
	return result
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
