package main

import (
	"fmt"
)

func uniqueOccurrences(arr []int) bool {
	m := make(map[int]int, len(arr))

	for _, v := range arr {
		m[v]++
	}

	unique := make(map[int]bool, len(m))
	for _, count := range m {
		if _, ok := unique[count]; ok {
			return false
		} else {
			unique[count] = true
		}
	}

	return true
}

func main() {
	arr := []int{1, 2, 2, 1, 1, 3}
	res := uniqueOccurrences(arr)
	fmt.Println(res)
}
