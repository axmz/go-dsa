package main

import (
	"fmt"
	"slices"
)

func splitArray(nums []int, k int) int {
	max := slices.Max(nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}

	left, right := max, sum
	for left <= right {
		mid := left + (right-left)/2
		if canSplit(nums, k, mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func canSplit(nums []int, k int, maxSum int) bool {
	countSplits := 0
	currentSum := 0
	for _, num := range nums {
		currentSum += num
		if currentSum > maxSum {
			countSplits++
			currentSum = num
			if countSplits >= k {
				return false
			}
		}
	}
	return true
}

func main() {
	m := 3
	nums := []int{7, 2, 5, 10, 8}
	fmt.Println(splitArray(nums, m))
}
