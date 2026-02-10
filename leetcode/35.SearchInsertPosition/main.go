package main

import "fmt"

func searchInsert(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if target < nums[i] {
			return i
		}
	}
	return len(nums)
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
