package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	return kSum(nums, target, 4)
}

func kSum(nums []int, target int, k int) [][]int {
	res := [][]int{}
	l := len(nums)

	if k == 2 {
		return twoSum(nums, target)
	}

	for i := 0; i < l-(k-1); i++ {
		if i == 0 || nums[i-1] != nums[i] {
			for _, subset := range kSum(nums[i+1:], target-nums[i], k-1) {
				res = append(res, append([]int{nums[i]}, subset...))
			}
		}
	}

	return res
}

func twoSum(nums []int, target int) [][]int {
	res := [][]int{}
	lo := 0
	hi := len(nums) - 1
	for lo < hi {
		curr_sum := nums[lo] + nums[hi]
		if curr_sum < target || (lo > 0 && nums[lo] == nums[lo-1]) {
			lo++
		} else if curr_sum > target || (hi < len(nums)-1 && nums[hi] == nums[hi+1]) {
			hi--
		} else {
			res = append(res, []int{nums[lo], nums[hi]})
			lo++
			hi--
		}
	}
	return res
}

func main() {
	target := 8
	nums := []int{2, 2, 2, 2, 2}
	// target := -1
	// nums := []int{1, 0, -1, 0, -2, 2}
	fmt.Println(fourSum(nums, target))
}
