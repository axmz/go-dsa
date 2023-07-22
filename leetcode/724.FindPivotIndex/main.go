package main

import (
	"fmt"
)

func pivotIndex(nums []int) int {
	sum := 0
	lsum := 0
	for _, v := range nums {
		sum += v
	}

	for i, v := range nums {
		if lsum == sum-nums[i] {
			return i
		} else {
			lsum += v
			sum -= v
		}
	}

	return -1
}

func main() {
	// nums := []int{1, 2, 3}
	// nums := []int{1, 7, 3, 6, 5, 6}
	// nums := []int{-1, -1, -1, -1, -1, 0}
	nums := []int{2, 1, -1}
	res := pivotIndex(nums)
	fmt.Println(res)
}
