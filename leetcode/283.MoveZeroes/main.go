package main

import (
	"fmt"
)

// This can be shortened
func moveZeroes(nums []int) {
	for i, j := 0, 0; i < len(nums); {
		if nums[j] == 0 && nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j++
			continue
		}

		if nums[j] != 0 {
			j++
		}

		i++
	}
}

func main() {
	a := []int{0, 1, 0, 3, 12}
	moveZeroes(a)
	fmt.Println(a)
}
