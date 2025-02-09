package main

import "fmt"

func sortedSquares(nums []int) []int {
	l := len(nums)
	res := make([]int, l)
	k := l - 1

	for i, j := 0, l-1; i <= j; {
		I := nums[i] * nums[i]
		J := nums[j] * nums[j]
		if I > J {
			res[k] = I
			i++
		} else {
			res[k] = J
			j--
		}
		k--
	}

	return res
}

func main() {
	nums := []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(nums))
}
