package main

import "fmt"

func checkIfExist(arr []int) bool {
	m := make(map[int]bool, len(arr))

	for _, v := range arr {
		if m[v*2] || (v%2 == 0 && m[v/2]) {
			return true
		}
		m[v] = true
	}
	return false
}

func main() {
	nums := []int{10, 2, 5, 3}
	// nums := []int{7, 1, 14, 11}
	// nums := []int{-2, 0, 10, -19, 4, 6, -8}
	fmt.Println(checkIfExist(nums))
}
