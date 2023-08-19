package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		dif := target - v
		idx, ok := m[dif]
		if ok {
			return []int{idx, i}
		}
		m[v] = i
	}
	return nil
}

func main() {
	fmt.Println("Two Sum:", twoSum([]int{2, 7, 11, 19}, 9))
}
