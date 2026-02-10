package main

import "fmt"

func evenNumberBitwiseORs(nums []int) int {
	sum := 0
	for _, n := range nums {
		if n&1 == 0 {
			sum += n
		}
	}
	return sum
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
