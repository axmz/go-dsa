package main

import "fmt"

func search(nums []int, l, r int) int {
	if l == r {
		return l
	}

	m := (l + r) / 2
	if nums[m] < nums[m+1] {
		return search(nums, m+1, r)
	} else {
		return search(nums, l, m)
	}

}

func rec(nums []int) int {
	l, r := 0, len(nums)-1
	return search(nums, l, r)
}

func iter(nums []int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		m := (l + r) / 2
		if nums[m] < nums[m+1] {
			l = m + 1
		} else {
			r = m
		}
	}

	return l
}

func findPeakElement(nums []int) int {
	return rec(nums)
}

func main() {
	// nums := []int{ 1,2,1,3,5,6,4 }
	nums := []int{1, 2, 3, 1}
	fmt.Println(findPeakElement(nums))
}
