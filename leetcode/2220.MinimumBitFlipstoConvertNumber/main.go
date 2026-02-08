package main

import "fmt"

func minBitFlips(start int, goal int) int {
	res := start ^ goal
	count := 0
	for res > 0 {
		count++
		res &= (res - 1)
	}
	return count
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
