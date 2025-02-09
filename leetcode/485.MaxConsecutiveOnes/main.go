package main

import "fmt"

func max(a, b int) int {

	if a > b {
		return a
	}

	return b
}

func findMaxConsecutiveOnes(nums []int) int {
	m := 0
	prev := 0
	consecutive := 0

	for _, v := range nums {
		if v == 1 {
			if prev == 1 {
				consecutive++
			} else {
				consecutive = 1
				prev = 1
			}
			m = max(m, consecutive)
		} else {
			prev = 0
		}
	}
	return m
}

func main() {
	nums := []int{1, 1, 0, 1, 1, 1}
	fmt.Println(findMaxConsecutiveOnes(nums))
}
