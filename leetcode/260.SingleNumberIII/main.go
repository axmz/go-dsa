package main

import (
	"fmt"
)

func singleNumber(nums []int) []int {
	xor := 0
	for _, n := range nums {
		xor ^= n
	}

	x := 0
	x_bitmap := xor & (-xor)

	for _, n := range nums {
		if x_bitmap&n != 0 {
			x ^= n
		}
	}

	return []int{x, xor ^ x}
}

func main() {
	nums := []int{1, 2, 1, 3, 2, 5}
	res := singleNumber(nums)
	fmt.Println(res)
}
