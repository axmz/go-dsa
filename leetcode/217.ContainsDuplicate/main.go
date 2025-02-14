package main

import "fmt"

func containsDuplicate(nums []int) bool {
	m := make(map[int]int)
	for _, v := range nums {
		if _, exists := m[v]; exists {
			return true
		}
		m[v]++
	}

	return false
}

func main() {
	nums := []int{}
	fmt.Println(containsDuplicate(nums))
}
