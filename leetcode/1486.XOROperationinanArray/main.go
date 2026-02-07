package main

import "fmt"

func xorOperation(n int, start int) int {
	xor := 0
	for i := 0; i < n; i++ {
		xor ^= start + 2*i
	}
	return xor
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
