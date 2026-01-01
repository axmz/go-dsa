package main

import "fmt"

func climb(n int, memo map[int]int) int {
	if n == 0 {
		return 1
	}

	if n < 0 {
		return 0
	}

	if val, ok := memo[n]; ok {
		return val
	}

	memo[n] = climb(n-1, memo) + climb(n-2, memo)
	return memo[n]
}

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	memo := make(map[int]int, n+1)
	return climb(n-1, memo) + climb(n-2, memo)
}

func main() {
	x := 6
	fmt.Println(climbStairs(x))
}
