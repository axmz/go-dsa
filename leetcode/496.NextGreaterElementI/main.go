package main

import "fmt"

func nextGreaterElement(nums1 []int, nums2 []int) []int {

	stack := make([]int, 0)
	var pop int
	res := make([]int, len(nums1))
	for i := range res {
		res[i] = -1
	}

	m := make(map[int]int, len(nums2))
	for i, v := range nums2 {
		if i == 0 {
			stack = append(stack, v)
			continue
		}

		for len(stack) > 0 && v >= stack[len(stack)-1] {
			pop, stack = stack[len(stack)-1], stack[:len(stack)-1]
			m[pop] = v
		}

		stack = append(stack, v)
	}

	for i, v := range nums1 {
		if t, exists := m[v]; exists {
			res[i] = t
		}
	}

	return res
}

func main() {

	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}
	// nums1 := []int{2, 4}
	// nums2 := []int{1, 2, 3, 4}
	r := nextGreaterElement(nums1, nums2)
	fmt.Println(r)
}
