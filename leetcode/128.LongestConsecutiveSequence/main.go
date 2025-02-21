package main

import "fmt"

func longestConsecutive(nums []int) int {
	m := make(map[int]bool, len(nums))

	for _, v := range nums {
		m[v] = true
	}

	max := 0
	for k := range m {
		if _, exists := m[k-1]; !exists {
			cur := 1
			i := 1
			for m[k+i] {
				cur++
				i++
			}
			if cur > max {
				max = cur
			}
		}
	}

	return max
}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(nums))
}
