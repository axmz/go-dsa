package main

import "fmt"

func prefixesDivBy5(nums []int) []bool {
	res := make([]bool, len(nums))
	n := 0
	for i, v := range nums {
		n = (n<<1 + v) % 5
		if n == 0 {
			res[i] = true
		}
	}
	return res
}

// The mistake in the prefixesDivBy5_2 function is that it does not reduce n modulo 5 at each step.
// As a result, n can become very large for long input arrays.
// This can lead to integer overflow or unnecessary computation.
// The correct approach is to keep n as the remainder modulo 5 after each bit is added,
// as done in the prefixesDivBy5 function.
func prefixesDivBy5_incorrect(nums []int) []bool {
	res := make([]bool, len(nums))
	n := 0
	for i, v := range nums {
		n <<= 1
		n |= v
		if n%5 == 0 {
			res[i] = true
		}
	}
	return res
}

func prefixesDivBy5_correct(nums []int) []bool {
	res := make([]bool, len(nums))
	n := 0
	for i, v := range nums {
		n = (n<<1 | v) % 5
		if n == 0 {
			res[i] = true
		}
	}
	return res
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
