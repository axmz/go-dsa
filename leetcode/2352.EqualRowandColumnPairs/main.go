package main

import (
	"fmt"
)

func equalPairs(grid [][]int) int {
	var s string
	n := len(grid)

	r := make(map[string]int, n)
	for _, v := range grid {
		s = fmt.Sprintf("%v", v)
		r[s] = r[s] + 1
	}

	sl := make([]int, n)
	c := make(map[string]int, n)
	for i := 0; i < n; i++ {
		sl = sl[:0]
		for j := 0; j < n; j++ {
			sl = append(sl, grid[j][i])
		}
		s = fmt.Sprintf("%v", sl)
		c[s] = c[s] + 1
	}

	fmt.Println(n, c)
	sum := 0
	for k, rv := range r {
		if cv, ok := c[k]; ok {
			sum += rv * cv
		}
	}

	return sum
}

func main() {
	grid := [][]int{{2, 2}, {2, 2}}
	// grid := [][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}}
	// grid := [][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}}
	a := equalPairs(grid)
	fmt.Println(a)

}
