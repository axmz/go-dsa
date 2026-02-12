package main

import "fmt"

func helper(nums []int) [][]int {
	res := [][]int{{nums[0]}}
	if len(nums) == 1 {
		return res
	}

	subsets := helper(nums[1:])
	for _, subset := range subsets {
		newSubset := append([]int{nums[0]}, subset...)
		res = append(res, subset)
		res = append(res, newSubset)
	}

	return res
}

func subsets(nums []int) [][]int {
	s := helper(nums)
	s = append(s, []int{})
	return s
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))
}
