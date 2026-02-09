package main

import "fmt"

func orArray(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i-1] |= nums[i]
	}
	return nums[:len(nums)-1]
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
