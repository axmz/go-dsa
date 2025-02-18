package main

import (
	"fmt"
	"sort"
)

func twoSum(nums []int, target int, res *[][]int) {
	l, r := 0, len(nums)-1

	for l < r {
		sum := nums[l] + nums[r]
		if sum == target {
			*res = append(*res, []int{-target, nums[l+1], nums[r+1]})
			l++
			r--
			for l < r && nums[l] == nums[l-1] {
				l++
			}
		}

		if sum > target {
			r--
		}

		if sum < target {
			l++
		}
	}
}

func threeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)

	for i := 0; i < len(nums) && nums[i] <= 0; i++ {
		if i == 0 || nums[i-1] != nums[i] {
			twoSum(nums[i+1:], -nums[i], &res)
		}
	}

	return res
}

func main() {
	nums := []int{}
	fmt.Println(threeSum(nums))
}
