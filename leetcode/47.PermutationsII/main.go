package main

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	res := [][]int{}
	l := len(nums)
	sort.Ints(nums)

	var permute func(nums []int, path []int)
	permute = func(nums []int, path []int) {
		if len(path) == l {
			res = append(res, path)
			return
		}
		for i, n := range nums {
			// skip duplicates
			if i > 0 && nums[i] == nums[i-1] {
				continue
			}

			remaining := append([]int{}, nums[:i]...)
			remaining = append(remaining, nums[i+1:]...)
			copyPath := append([]int{}, path...)
			copyPath = append(copyPath, n)
			permute(remaining, copyPath)
		}
	}

	permute(nums, []int{})
	return res
}

func permuteUniqueOptimized(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	used := make([]bool, len(nums))

	var permute func(path []int)
	permute = func(path []int) {
		if len(path) == len(nums) {
			res = append(res, append([]int{}, path...))
			return
		}
		for i := 0; i < len(nums); i++ {
			// skip if already used or duplicate
			if used[i] || (i > 0 && nums[i] == nums[i-1] && !used[i-1]) {
				continue
			}
			used[i] = true
			permute(append(path, nums[i]))
			used[i] = false
		}
	}

	permute([]int{})
	return res
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
