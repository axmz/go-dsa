package main

import "fmt"

func fourSumCount2(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	l := len(nums1)
	m := make(map[int]int, l)
	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			m[v1+v2]++
		}
	}

	res := 0
	for _, v3 := range nums3 {
		for _, v4 := range nums4 {
			res += m[-v3-v4]
		}
	}

	return res
}

func fourSumCount(A, B, C, D []int) int {
	lsts := [][]int{A, B, C, D}
	k := len(lsts)

	left := sumCount(lsts, 0, k/2)
	right := sumCount(lsts, k/2, k)

	res := 0
	for s, count := range left {
		res += count * right[-s]
	}
	return res
}

func sumCount(lsts [][]int, start, end int) map[int]int {
	cnt := map[int]int{0: 1}

	for i := start; i < end; i++ {
		newMap := make(map[int]int)
		for _, num := range lsts[i] {
			for total, count := range cnt {
				newMap[total+num] += count
			}
		}
		cnt = newMap
	}
	return cnt
}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{-2, -1}
	nums3 := []int{-1, 2}
	nums4 := []int{0, 2}
	fmt.Println(fourSumCount(nums1, nums2, nums3, nums4))
}
