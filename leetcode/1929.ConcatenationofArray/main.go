package main

import "fmt"

func getConcatenation(nums []int) []int {
	n := len(nums)
	ans := make([]int, 2*n)
	for i := range ans {
		ans[i] = nums[i%n]
	}
	return ans
}

func main() {
	// nums := []int{1, 2, 3}
	nums := []int{}
	fmt.Println(getConcatenation(nums))
}
