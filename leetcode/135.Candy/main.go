package main

import "fmt"

func candy(ratings []int) int {
	l := len(ratings)
	leftToRight := make([]int, l)
	rightToLeft := make([]int, l)

	for i := 1; i < l; i++ {
		if ratings[i] > ratings[i-1] {
			leftToRight[i] = leftToRight[i-1] + 1
		}
	}

	for i := l - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			rightToLeft[i] = rightToLeft[i+1] + 1
		}
	}

	candies := 0
	for i := 0; i < l; i++ {
		candies += max(leftToRight[i], rightToLeft[i]) + 1
	}

	return candies
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
