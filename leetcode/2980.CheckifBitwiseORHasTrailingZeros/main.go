package main

import "fmt"

func hasTrailingZeros(nums []int) bool {
	countEven := 0
	for _, n := range nums {
		if n&1 == 0 {
			countEven++
			if countEven >= 2 {
				return true
			}
		}
	}
	return false
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
