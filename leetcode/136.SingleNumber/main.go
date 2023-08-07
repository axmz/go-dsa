package main

import (
	"fmt"
)

func singleNumber(nums []int) int {
	xor := 0
	for _, v := range nums {
		xor ^= v
		fmt.Printf("%d - %08b - %08b \n", v, v, xor)
	}
	return xor
}

func main() {
	nums := []int{4, 1, 2, 1, 2, 4, 4, 1, 4}
	res := singleNumber(nums)
	fmt.Println(res)
}
