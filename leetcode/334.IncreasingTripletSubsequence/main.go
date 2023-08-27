package main

import (
	"fmt"
	"math"
)

func increasingTriplet(nums []int) bool {
	maxInt32 := math.MaxInt32
	lowest1, lowest2, lowest3 := maxInt32, maxInt32, maxInt32

	for _, v := range nums {
		switch true {
		case v <= lowest1:
			lowest1 = v
		case v <= lowest2:
			lowest2 = v
		case v <= lowest3:
			return true
		}

	}
	return false
}

func main() {
	// nums := []int{1, 2, 3, 4, 5}
	// nums := []int{1, 1, 1, 1, 1}
	nums := []int{4, 5, 2147483647, 1, 2}
	// nums := []int{5, 4, 3, 2, 1}
	// nums := []int{2, 1, 5, 0, 4, 6}
	res := increasingTriplet(nums)
	fmt.Println(res)
}
