package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var res [][]int
	var backtrack func(start int, path []int, sum int)
	backtrack = func(start int, path []int, sum int) {
		if sum == target {
			res = append(res, path)
			return
		}

		if sum > target {
			return
		}

		for i, n := range candidates[start:] {
			// this is clever
			if i > 0 && n == candidates[start+i-1] {
				continue
			}
			p := append([]int{}, path...)
			p = append(p, n)
			s := sum + n
			backtrack(start+i+1, p, s)
		}
	}

	backtrack(0, []int{}, 0)

	return res
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
