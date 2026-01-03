package main

import "fmt"

func maximumWealth(accounts [][]int) int {
	max := 0
	for _, a := range accounts {
		sum := 0
		for _, v := range a {
			sum += v
		}
		if sum > max {
			max = sum
		}
	}
	return max
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
