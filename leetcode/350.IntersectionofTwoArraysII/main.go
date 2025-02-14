package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int, len(nums1))
	res := make([]int, 0)
	for _, v := range nums1 {
		m[v]++
	}
	for _, v := range nums2 {

		if val, exists := m[v]; exists && val > 0 {
			res = append(res, v)
			m[v]--
		}
	}

	return res

}

func main() {
	nums1, nums2 := []int{4, 9}, []int{9, 4}
	fmt.Println(intersect(nums1, nums2))
}
