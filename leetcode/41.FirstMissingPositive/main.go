package main

import (
	"fmt"
)

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

func firstMissingPositive2(nums []int) int {
	l := len(nums)

	for i, n := range nums {
		if n < 1 || n > l {
			nums[i] = l + 1
		}
	}

	for _, n := range nums {
		n = abs(n)
		if n > 0 && n <= l {
			nums[n-1] = -abs(nums[n-1])
		}
	}

	for i, n := range nums {
		if n > 0 {
			return i + 1
		}
	}

	return l + 1
}
func firstMissingPositive(nums []int) int {
	l := len(nums)

	for i, n := range nums {
		if n < 1 || n > l {
			nums[i] = l + 1
		}
	}

	for i := 0; i < l; i++ {
		for nums[i]-1 < l && nums[i] != nums[nums[i]-1] {
			j := nums[i] - 1
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	for i := 0; i < l; i++ {
		if nums[i]-1 != i {
			return i + 1
		}
	}

	return l + 1
}

func main() {
	// nums := []int{1, 1}
	// nums := []int{3, 4, -1, 1}
	nums := []int{1, 3, 2}
	fmt.Println(firstMissingPositive(nums))
}
