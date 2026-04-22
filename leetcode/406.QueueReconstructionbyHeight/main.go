package main

import (
	"sort"
)

func reconstructQueue(people [][]int) [][]int {
	// sort by height desc, then by k asc
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})

	res := make([][]int, 0, len(people))
	for i := range people {
		k := people[i][1]
		// linked list would be great here
		res = append(res[:k], append([][]int{people[i]}, res[k:]...)...)
	}

	return res
}

func main() {
	people := [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}
	reconstructQueue(people)
}
