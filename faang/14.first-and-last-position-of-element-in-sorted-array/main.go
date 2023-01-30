package main

import (
	"fmt"
)

func assign(i int, left, right *int) {
	if i < *left || *left == -1 {
		*left = i
	}
	if i > *right || *right == -1 {
		*right = i
	}
}

func search(nums []int, target, i, j int, left, right *int) {
	if i == j && nums[i] == target {
		assign(i, left, right)
	}

	if i > j {
		return
	}

	l := j - i
	m := i + l/2
	t := nums[m]

	if t < target {
		search(nums, target, m+1, j, left, right)
	} else if t > target {
		search(nums, target, i, m-1, left, right)
	} else {
		assign(m, left, right)
		search(nums, target, i, m-1, left, right)
		search(nums, target, m+1, j, left, right)
	}
}

func searchRange(nums []int, target int) []int {
	var left, right = -1, -1
	search(nums, target, 0, len(nums)-1, &left, &right)
	return []int{left, right}
}

func main() {
	// nums := []int{}
	// nums := []int{5, 7, 7, 8, 8, 10}
	nums := []int{1, 2, 2, 3, 5, 5, 5, 5, 5, 6, 7}
	r := searchRange(nums, 5)
	fmt.Println(r)
}
