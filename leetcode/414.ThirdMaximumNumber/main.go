package main

import (
	"fmt"
	"math"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func thirdMax(nums []int) int {
	max1, max2, max3 := math.MinInt, math.MinInt, math.MinInt

	for _, v := range nums {
		max1 = max(max1, v)
	}

	for _, v := range nums {
		if v == max1 {
			continue
		}
		max2 = max(max2, v)
	}
	if max2 == math.MinInt {
		return max1
	}

	for _, v := range nums {
		if v == max1 || v == max2 {
			continue
		}
		max3 = max(max3, v)
	}
	if max3 == math.MinInt {
		return max1
	}

	return max3
}

func main() {
	// nums := []int{3, 2, 1, 4}
	nums := []int{1, 2}
	// nums := []int{2,2,3,1}
	fmt.Println(thirdMax(nums))
}
