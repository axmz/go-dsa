package main

import "fmt"

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	return n&(n-1) == 0
}

func isPowerOfTwo2(n int) bool {
	if n <= 0 {
		return false
	}
	return n&(-n) == n
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
