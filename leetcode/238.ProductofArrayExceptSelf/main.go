package main

import (
	"fmt"
)

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	for i := range res {
		res[i] = 1
	}

	left := 1
	right := 1

	for i, j := 1, len(nums)-2; i < len(nums); i, j = i+1, j-1 {
		left *= nums[i-1]
		res[i] *= left
		right *= nums[j+1]
		res[j] *= right
	}

	return res
}

func main() {

	nums := []int{1, 2, 3, 4}
	res := productExceptSelf(nums)
	fmt.Println(res)
}
