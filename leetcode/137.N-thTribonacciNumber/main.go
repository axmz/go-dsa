package main

import (
	"fmt"
)

func tribonacci3(n int) int {
	memo := [3]int{0, 1, 1}
	if n < 3 {
		return memo[n]
	}
	for i := 3; i <= n; i++ {
		memo[0], memo[1], memo[2] = memo[1], memo[2], memo[0]+memo[1]+memo[2]
	}
	return memo[2]
}

func tribonacci2(n int) int {
	if n < 3 {
		if n == 0 {
			return 0
		} else {
			return 1
		}
	}

	t := make([]int, n+1)
	t[0] = 0
	t[1] = 1
	t[2] = 1
	for i := 3; i <= n; i++ {
		t[i] = t[i-1] + t[i-2] + t[i-3]
	}

	return t[n]
}

func tribonacci(n int) int {
	if n < 3 {
		if n == 0 {
			return 0
		} else {
			return 1
		}
	}

	var a, b, c, sum int = 0, 1, 1, 0

	for i := 3; i <= n; i++ {
		sum = a + b + c
		a = b
		b = c
		c = sum
	}

	return sum
}

func main() {
	res := tribonacci(25)
	fmt.Println(res)
}
