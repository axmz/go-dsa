package main

import "fmt"

func longestSubarray(nums []int) int {
	var l, r, max, newMax, count, lastL int

	for l, r = 0, 0; r < len(nums); r++ {
		if nums[r] == 0 {
			if count == 0 {
				count++
				lastL = r
			} else {
				l = lastL + 1
				lastL = r
			}
		}

		newMax = r + 1 - l - 1
		if newMax > max {
			max = newMax
		}
	}

	return max
}

func main() {

	// nums := []int{1} // 0
	nums := []int{0} // 0
	// nums := []int{0, 0} // 0
	// nums := []int{1, 1} // 1
	// nums := []int{1, 1, 0, 1} // 3
	// nums := []int{0, 1, 1, 1, 0, 1, 1, 0, 1} // 5
	// nums := []int{0, 0, 1, 1, 1, 0, 0, 1, 1, 0, 1} // 3
	// nums := []int{1, 1, 1} // 2
	// nums := []int{1, 1, 0, 0, 1, 1, 1, 0, 1} // 4

	res := longestSubarray(nums)
	fmt.Println(res)
}
