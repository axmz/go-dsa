package main

import (
	"fmt"
	"math"
)

func minSubArrayLen(target int, nums []int) int {
	sum := 0
	min := math.MaxInt

	for r, l := 0, 0; r < len(nums); r++ {
		sum += nums[r]
		for sum >= target {
			if r-l < min {
				min = r - l
			}
			sum -= nums[l]
			l++
		}
	}

	if min == math.MaxInt {
		return 0
	}

	return min + 1
}

func main() {
	target := 4
	nums := []int{1, 4, 4}
	// target := 7
	// nums := []int{2, 3, 1, 2, 4, 3}
	fmt.Println(minSubArrayLen(target, nums))
}
