package main

import (
	"fmt"
	"sort"
)

func makesquare(matchsticks []int) bool {
	if len(matchsticks) < 4 {
		return false
	}

	sum := 0
	for _, m := range matchsticks {
		sum += m
	}
	if sum%4 != 0 {
		return false
	}

	side := sum / 4

	sort.Sort(sort.Reverse(sort.IntSlice(matchsticks)))

	sides := make([]int, 4)
	var backtrack func(int) bool
	backtrack = func(index int) bool {
		if index == len(matchsticks) {
			return sides[0] == sides[1] && sides[1] == sides[2] && sides[2] == sides[3]
		}
		for i := 0; i < 4; i++ {
			if sides[i]+matchsticks[index] <= side {
				sides[i] += matchsticks[index]
				if backtrack(index + 1) {
					return true
				}
				sides[i] -= matchsticks[index]
			}
		}
		return false
	}

	return backtrack(0)
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
