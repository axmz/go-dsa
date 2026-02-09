package main

import "fmt"

func duplicateNumbersXOR(nums []int) int {
	bitset := 0
	xor := 0
	for _, n := range nums {
		bit := bitset & (1 << n)
		if bit == 0 {
			bitset |= 1 << n
		} else {
			xor ^= n
		}
	}
	return xor
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
