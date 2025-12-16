package main

import "fmt"

func sumOfSquares(nums []int) int {
	l := len(nums)
	sum := 0

	for i, v := range nums {
		if l%(i+1) == 0 {
			sum += v * v
		}
	}

	return sum
}

func main() {
	// x := []int{1, 2, 3, 4}
	x := []int{2, 7, 1, 19, 18, 3}
	fmt.Println(x, sumOfSquares(x))
}
