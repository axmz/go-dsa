package main

import "fmt"

func dominantIndex(nums []int) int {
	max := 0
	dmax := 0
	maxIdx := -1

	for i, v := range nums {
		if v > max {
			max = v
			if v >= dmax {
				maxIdx = i
			} else {
				maxIdx = -1
			}
			dmax = 2 * v
		} else {
			if v*2 > max {
				maxIdx = -1
			}
		}

	}

	return maxIdx
}

func main() {
	// nums := []int{3, 6, 1, 0}
	// nums := []int{1, 2, 3, 4}
	nums := []int{0, 0, 3, 2}
	fmt.Println(dominantIndex(nums))
}
