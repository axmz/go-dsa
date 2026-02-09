package main

import "fmt"

func minOperations2(nums []int, k int) int {
	target := 1
	for i := 0; i < k; i++ {
		target |= 1 << i
	}

	operations := 0
	for i := 0; i < len(nums); i++ {
		operations++
		n := nums[len(nums)-1-i] - 1
		target &= ^(1 << n)
		if target == 1 {
			break
		}
	}

	return operations
}

func minOperations(nums []int, k int) int {
	target := 1
	for i := 0; i < k; i++ {
		target |= 1 << i
	}

	operations := 0
	collection := 0
	for i := 0; i < len(nums); i++ {
		n := nums[len(nums)-1-i] - 1
		collection |= 1 << n
		operations++
		if collection&target == target {
			break
		}
	}

	return operations
}

func main() {
	k := 3
	nums := []int{3, 2, 5, 3, 1}
	fmt.Println(minOperations(nums, k))
}
