package main

import (
	"fmt"
	"math/bits"
)

func minFlips(a int, b int, c int) int {
	xor := (a | b) ^ c
	return bits.OnesCount(uint(xor)) + bits.OnesCount(uint(a&b&xor))
}

func main() {
	a, b, c := 2, 6, 5
	fmt.Println(minFlips(a, b, c))
	// 10110110
	// 01010101
	// 10101010
}
