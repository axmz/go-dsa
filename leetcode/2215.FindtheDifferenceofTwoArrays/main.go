package main

import (
	"fmt"
)

func findDifference(nums1 []int, nums2 []int) [][]int {
	m := make(map[int]bool, len(nums2))
	n1 := make(map[int]int, len(nums1))
	n2 := make(map[int]int, len(nums2))

	for _, v := range nums2 {
		if _, ok := m[v]; !ok {
			m[v] = false
		}
	}

	for _, v := range nums1 {
		if _, ok := m[v]; !ok {
			n1[v] = v
		} else {
			m[v] = true
		}
	}

	for _, v := range nums2 {
		if seen := m[v]; !seen {
			n2[v] = v
		}
	}

	r1 := []int{}
	for k := range n1 {
		r1 = append(r1, k)
	}
	r2 := []int{}
	for k := range n2 {
		r2 = append(r2, k)
	}
	return [][]int{r1, r2}
}

func main() {
	nums1 := []int{1, 2, 3}
	nums2 := []int{2, 4, 6}
	res := findDifference(nums1, nums2)
	fmt.Println(res)
}
