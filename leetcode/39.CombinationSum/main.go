package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
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

		for i := start; i < len(candidates); i++ {
			n := candidates[i]
			p := append([]int{}, path...)
			p = append(p, n)
			s := sum + n
			backtrack(i, p, s)
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
