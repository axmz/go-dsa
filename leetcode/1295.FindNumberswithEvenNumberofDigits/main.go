package main

import "fmt"

func findNumbers(nums []int) int {
	res := 0
	for _, n := range nums {
		counter := 0
		div := n
		for div != 0 {
			div /= 10
			counter++
		}
		if counter%2 == 0 {
			res++
		}
	}

	return res
}

func main() {
	nums := []int{12, 345, 2, 6, 7896}
	fmt.Println(findNumbers(nums))
}
