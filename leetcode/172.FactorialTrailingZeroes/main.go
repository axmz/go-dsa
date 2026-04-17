package main

import "fmt"

func trailingZeroes(n int) int {
	zeros := 0
	for i := 5; i <= n; i += 5 {
		cur := i
		for cur%5 == 0 {
			zeros++
			cur /= 5
		}
	}
	return zeros
}

func trailingZeroes2(n int) int {
	zeroCount := 0
	for i := 5; i <= n; i += 5 {
		powerOf5 := 5
		for i%powerOf5 == 0 {
			zeroCount += 1
			powerOf5 *= 5
		}
	}
	return zeroCount
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
