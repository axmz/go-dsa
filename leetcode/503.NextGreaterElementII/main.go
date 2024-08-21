package main

import "fmt"

func nextGreaterElements(nums []int) []int {
	stack := make([]int, 0)
	var pop int
	res := make([]int, len(nums))

	for i := range res {
		res[i] = -1
	}

	for i := 0; i < 2*len(nums); i++ {
		i := i % len(nums)
		v := nums[i]

		for len(stack) > 0 && v > nums[stack[len(stack)-1]] {
			pop, stack = stack[len(stack)-1], stack[:len(stack)-1]
			res[pop] = v
		}

		stack = append(stack, i)
	}

	return res
}

func main() {
	// nums := []int{1, 2, 1}
	nums := []int{6, 4, 5, 4, 3, 2, 1}
	// nums := []int{1, 2, 3, 4, 3}
	// nums := []int{100, 1, 11, 1, 120, 111, 123, 1, -1, -100}
	r := nextGreaterElements(nums)
	fmt.Println(r)
}
