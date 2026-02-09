package main

import "fmt"

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maximumStrongPairXor(nums []int) int {
	maxXor := 0

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			a := nums[i]
			b := nums[j]
			if abs(a-b) <= min(a, b) {
				xor := a ^ b
				if xor > maxXor {
					maxXor = xor
				}
			}
		}
	}

	return maxXor
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
