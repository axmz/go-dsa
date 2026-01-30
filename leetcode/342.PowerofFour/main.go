package main

import "fmt"

// There are other solutions as well
// using log and mod
// https://leetcode.com/problems/power-of-four/editorial/?envType=problem-list-v2&envId=bit-manipulation#approach-2-math
func isPowerOfFour2(n int) bool {
	if n <= 0 {
		return false
	}

	for i := 0; i < 32; i += 2 {
		if n == 1<<i {
			return true
		}
	}

	return false
}

func isPowerOfFour3(n int) bool {
	powerOf4Mask := 0x55555555 // binary: ...0101
	return n > 0 && n&(n-1) == 0 && n&powerOf4Mask != 0
}
func isPowerOfFour(n int) bool {
	powerOf4Mask := 0xAAAAAAAA // binary: ...1010
	return n > 0 && n&(n-1) == 0 && n&powerOf4Mask == 0
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
