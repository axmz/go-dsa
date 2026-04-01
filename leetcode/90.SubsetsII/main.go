package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	var backtrack func(start int, path []int)
	backtrack = func(start int, path []int) {
		res = append(res, path)

		for i, n := range nums[start:] {
			if i > 0 && n == nums[start+i-1] {
				continue
			}
			p := append([]int{}, path...)
			p = append(p, n)
			backtrack(start+i+1, p)
		}
	}

	backtrack(0, []int{})

	return res
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
