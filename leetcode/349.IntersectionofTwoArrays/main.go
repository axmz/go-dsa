// multiple other solutions: hashmap, two pointers after sorting
// but it is simple anyways
// yet keep in mind the sets as posibility for similat problems
package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	seen := make([]int, 1001)
	for i := range nums1 {
		seen[nums1[i]]++
	}

	res := make([]int, 0)
	for i := range nums2 {
		if seen[nums2[i]] > 0 {
			res = append(res, nums2[i])
			seen[nums2[i]] = 0
		}
	}

	return res
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
