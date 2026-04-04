package main

import "fmt"

func jump(nums []int) int {
	jumps := 0
	farthest := 0
	currentEnd := 0

	for i := 0; i < len(nums)-1; i++ {
		n := nums[i]
		if farthest < i+n {
			farthest = i + n
		}
		if i == currentEnd {
			jumps++
			currentEnd = farthest
		}
	}

	return jumps
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
