package main

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
