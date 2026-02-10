package main

import (
	"fmt"
	"math/bits"
)

func kthCharacter(k int) byte {
	return byte(bits.OnesCount(uint(k))%26) + 'a'
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
