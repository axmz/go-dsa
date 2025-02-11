package main

import (
	"fmt"
)

func findDisappearedNumbers2(nums []int) []int {
	m := make([]int, len(nums)+1)
	for _, v := range nums {
		m[v] = 1
	}

	res := []int{}
	for i := 1; i <= len(nums); i++ {
		if m[i] == 0 {
			res = append(res, i)
		}
	}

	return res
}

func abs(a int) int {
	if a >= 0 {
		return a
	}

	return -a
}

func findDisappearedNumbers(nums []int) []int {
	for _, v := range nums {
		v = abs(v)
		nums[v-1] = -abs(nums[v-1])
	}

	res := []int{}
	for i, v := range nums {
		if v > 0 {
			res = append(res, i+1)
		}
	}

	return res
}

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDisappearedNumbers(nums))
}
