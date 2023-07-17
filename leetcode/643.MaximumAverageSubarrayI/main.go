package main

import (
	"fmt"
)

func findMaxAverage(nums []int, k int) float64 {
	var sum float64 = 0
	for i := 0; i < k; i++ {
		sum += float64(nums[i])
	}
	var max = sum

	for i := k; i < len(nums); i++ {
		sum = sum + float64(nums[i]) - float64(nums[i-k])
		if sum > max {
			max = sum
		}
	}

	return max / float64(k)
}

// there are solutions that utilise one for loop to first calculate the sum and then slide the window
func main() {
	// nums := []int{1, 12, -5, -6, 50, 3}
	// k := 4
	nums := []int{4, 2, 1, 3, 3}
	k := 2
	res := findMaxAverage(nums, k)
	fmt.Println(res)
}
