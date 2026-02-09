package main

import (
	"fmt"
	"math/bits"
)

func minChanges(n int, k int) int {
	if n&k^k == 0 {
		return bits.OnesCount(uint(n ^ k))
	} else {
		return -1
	}
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
