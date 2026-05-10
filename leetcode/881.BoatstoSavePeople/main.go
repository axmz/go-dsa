package main

import (
	"fmt"
	"sort"
)

func numRescueBoats(people []int, limit int) int {
	l, r := 0, len(people)-1
	sort.Ints(people)
	boats := 0
	for l <= r {
		if people[l]+people[r] <= limit {
			l++
		}
		r--
		boats++
	}
	return boats
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
