package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func findMaxConsecutiveOnes2(nums []int) int {
	m := 0
	lastZeroIdx := 0
	hearts := 1
	count := 0

	for i := 0; i < len(nums); {
		if nums[i] == 1 {
			count++
			i++
		} else if nums[i] == 0 && hearts > 0 {
			count++
			hearts--
			lastZeroIdx = i
			i++
		} else {
			m = max(m, count)
			count = 0
			hearts = 1
			i = lastZeroIdx + 1
		}
		m = max(m, count)
	}

	return m
}

func findMaxConsecutiveOnes(nums []int) int {
	l, r := 0, 0
	usedZeros := 0
	m := 0

	for r < len(nums) {
		if nums[r] == 0 {
			usedZeros++
		}

		for usedZeros == 2 {
			if nums[l] == 0 {
				usedZeros--
			}
			l++
		}

		r++

		m = max(m, r-l)
	}

	return m
}

func main() {
	nums := []int{1, 0, 1, 1, 0, 1}
	// nums := []int{1, 0, 1, 1, 0}
	fmt.Println(findMaxConsecutiveOnes(nums))
}
