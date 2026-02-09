package main

import (
	"fmt"
	"math/bits"
)

func tripletCount(a []int, b []int, c []int) int {
	count := 0
	for _, x := range a {
		for _, y := range b {
			for _, z := range c {
				xor := x ^ y ^ z
				if bits := bits.OnesCount(uint(xor)); bits%2 == 0 {
					count++
				}
			}
		}
	}
	return count
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
