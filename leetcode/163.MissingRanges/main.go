package main

import "fmt"

func findMissingRanges(nums []int, lower int, upper int) [][]int {
	res := [][]int{}

	if len(nums) == 0 {
		res = append(res, []int{lower, upper})
		return res
	}

	for _, v := range nums {
		if v-lower > 0 {
			res = append(res, []int{lower, v - 1})
		}
		lower = v + 1
	}

	if upper > nums[len(nums)-1] {
		res = append(res, []int{lower, upper})
	}

	return res
}

func main() {
	lower := 0
	upper := 99
	nums := []int{0, 1, 3, 50, 75}
	fmt.Println(findMissingRanges(nums, lower, upper))
}
