package main

import "fmt"

func findKOr(nums []int, k int) int {
	kor := 0

	for i := 0; i < 32; i++ {
		sum := 0
		for _, num := range nums {
			sum += 1 & (num >> i)
			if sum == k {
				kor |= 1 << i
				break
			}
		}
	}

	return kor
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
