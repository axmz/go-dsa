package main

import (
	"fmt"
)

func moveZeroes(nums []int) {
	for i, j := 0, 0; j < len(nums); j++ {
		if nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
}

func main() {
	a := []int{0, 1, 0, 3, 12}
	moveZeroes(a)
	fmt.Println(a)
}
