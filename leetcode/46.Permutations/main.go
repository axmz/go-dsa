package main

import "fmt"

func permute(nums []int) [][]int {
	res := [][]int{}

	var backtrack func(path []int, used []bool)
	backtrack = func(path []int, used []bool) {
		if len(path) == len(nums) {
			// make a copy of path
			perm := make([]int, len(path))
			copy(perm, path)
			res = append(res, perm)
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			// make choice
			used[i] = true
			path = append(path, nums[i])
			backtrack(path, used)
			// unmake choice
			used[i] = false
			path = path[:len(path)-1]
		}

	}

	backtrack([]int{}, make([]bool, len(nums)))
	return res
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
