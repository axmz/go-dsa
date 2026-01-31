package main

import (
	"fmt"
	"math"
)

func findComplement2(num int) int {
	mask := 1
	n := num
	for n > 0 {
		n >>= 1
		mask <<= 1
	}
	return ^num & (mask - 1)
}

func findComplement3(num int) int {
	mask := 0
	n := num
	for n > 0 {
		n >>= 1
		mask = (mask << 1) | 1
	}
	return ^num & mask
}

func msb(num int) int {
	mask := num
	mask |= mask >> 1
	mask |= mask >> 2
	mask |= mask >> 4
	mask |= mask >> 8
	mask |= mask >> 16
	return mask - (mask >> 1)
}

func findComplement4(num int) int {
	return ^num & (msb(num)<<1 - 1)
}

func findComplement5(num int) int {
	mask := num
	mask |= mask >> 1
	mask |= mask >> 2
	mask |= mask >> 4
	mask |= mask >> 8
	mask |= mask >> 16
	return ^num & mask
}

func findComplement(num int) int {
	l := math.Log2(float64(num)) + 1
	mask := (1 << uint(l)) - 1
	return num ^ mask // or return ^num & mask
}

func main() {
	fmt.Println(findComplement(5))
}
