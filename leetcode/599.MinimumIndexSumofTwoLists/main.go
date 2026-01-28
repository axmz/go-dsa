package main

import "fmt"

// Easy but the solution is not trivial, rather elegant
func findRestaurant(list1 []string, list2 []string) []string {
	m := make(map[string]int, len(list1))
	for i, v := range list1 {
		m[v] = i
	}
	minSum := len(list1) + len(list2) // nice technique to set initial max value
	res := []string{}
	for j, v := range list2 {
		if i, ok := m[v]; ok {
			if i+j < minSum { // tracking min and updating result in one go not two
				minSum = i + j
				res = []string{v}
			} else if i+j == minSum {
				res = append(res, v)
			}
		}
	}
	return res
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
