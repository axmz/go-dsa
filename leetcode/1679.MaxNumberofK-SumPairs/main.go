package main

import (
	"fmt"
)

func maxOperations(nums []int, k int) int {
	count := 0
	pair := make(map[int]int)
	for _, v := range nums {
		dif := k - v
		if p := pair[v]; p > 0 {
			count++
			pair[v]--
		} else {
			pair[dif]++
		}
	}
	return count
}

func main() {
	nums := []int{1, 2, 3, 4}
	k := 5
	// nums := []int{3, 1, 3, 4, 3}
	// k := 6
	// nums := []int{2, 2, 2, 3, 1, 1, 4, 1}
	// k := 4
	res := maxOperations(nums, k)
	fmt.Println(res)
}
