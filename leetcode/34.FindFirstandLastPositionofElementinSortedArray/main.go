package main

import (
	"fmt"
)

func search_right(nums []int, l, r, target int) int {
	if l > r {
		return -1
	}

	mid := l + (r-l)/2
	if nums[mid] == target {
		right := search_right(nums, mid+1, r, target)
		if right == -1 {
			return mid
		}
		return right
	} else if nums[mid] < target {
		return search_right(nums, mid+1, r, target)
	} else {
		return search_right(nums, l, mid-1, target)
	}
}

func search_left(nums []int, l, r, target int) int {
	if l > r {
		return -1
	}

	mid := l + (r-l)/2
	if nums[mid] == target {
		left := search_left(nums, l, mid-1, target)
		if left == -1 {
			return mid
		}
		return left
	} else if nums[mid] < target {
		return search_left(nums, mid+1, r, target)
	} else {
		return search_left(nums, l, mid-1, target)
	}
}

func search_left2(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	if l < len(nums) && nums[l] == target {
		return l
	}
	return -1
}

func search_right2(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			l = mid + 1
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	if r >= 0 && nums[r] == target {
		return r
	}
	return -1
}

func searchRange(nums []int, target int) []int {
	// left := search_left(nums, 0, len(nums)-1, target)
	// right := search_right(nums, 0, len(nums)-1, target)
	left := search_left2(nums, target)
	right := search_right2(nums, target)
	return []int{left, right}
}

func main() {
	// nums := []int{}
	// nums := []int{5, 7, 7, 8, 8, 10}
	nums := []int{1, 2, 2, 3, 5, 5, 5, 5, 5, 6, 7}
	r := searchRange(nums, 5)
	fmt.Println(r)
}
