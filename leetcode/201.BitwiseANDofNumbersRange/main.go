package main

import "fmt"

func rangeBitwiseAnd(left int, right int) int {
	for left < right {
		// turn off rightmost 1-bit
		right = right & (right - 1)
	}

	return right
}

func rangeBitwiseAnd2(left int, right int) int {
	var shift int

	for left != right {
		left >>= 1
		right >>= 1
		shift++
	}

	return left << shift
}

func main() {
	left, right := 4, 9
	fmt.Println(rangeBitwiseAnd(left, right))
}
