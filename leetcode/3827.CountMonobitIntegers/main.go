package main

import "fmt"

func countMonobit(n int) int {
	count := 0
	monobit := 1
	for monobit <= n {
		count++
		monobit = monobit<<1 | 1
	}
	return count + 1
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
