package main

import "fmt"

func numberOfSteps(num int) int {
	steps := 0
	for num > 0 {
		if num%2 == 0 {
			num >>= 1
		} else {
			num -= 1
		}
		steps++
	}

	return steps
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
