package main

import "fmt"

func wiggleSort(nums []int) {
	n := len(nums)
	for i := 1; i < n; i++ {
		// if i % 2 == 1 {
		// 	if nums[i] < nums[i-1] {
		// 		nums[i], nums[i-1] = nums[i-1], nums[i]
		// 	}
		// } else {
		// 	if nums[i] > nums[i-1] {
		// 		nums[i], nums[i-1] = nums[i-1], nums[i]
		// 	}
		// }
		if i%2 == 1 && nums[i] < nums[i-1] || i%2 == 0 && nums[i] > nums[i-1] {
			nums[i], nums[i-1] = nums[i-1], nums[i]
		}
	}
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
