package main

import "fmt"

func findDuplicate(nums []int) int {
	l := len(nums)
	h := nums[0]
	t := nums[0]

	for {
		h = nums[nums[h]]
		t = nums[t]
		if h >= l {
			return -1
		}

		if h == t {
			h = nums[0]
			break
		}
	}

	for h != t {
		h = nums[h]
		t = nums[t]
	}

	return h
}

func main() {
	nums := []int{1, 3, 4, 2, 2}
	fmt.Println(findDuplicate(nums))
}
