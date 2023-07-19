package main

import (
	"fmt"
)

func max(a, b int) int {

	if a > b {
		return a
	}
	return b
}

func largestAltitude(gain []int) int {
	prev := 0
	maxAlt := prev
	for _, v := range gain {
		prev += v
		if prev > maxAlt {
			maxAlt = prev
		}
	}

	return maxAlt
}

func main() {
	// gain := []int{-5, 1, 5, 0, -7}
	gain := []int{-4, -3, -2, -1, 4, 3, 2}
	res := largestAltitude(gain)
	fmt.Println(res)
}
