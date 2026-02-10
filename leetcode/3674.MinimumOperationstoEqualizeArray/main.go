package main

import "fmt"

func minOperations(nums []int) int {
	for i, n := range nums[1:] {
		if n != nums[i] {
			return 1
		}
	}
	return 0
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
