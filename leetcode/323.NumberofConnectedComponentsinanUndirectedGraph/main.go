package main

import (
	"fmt"
	. "godsa/utils/disjointset"
)

func countComponents(n int, edges [][]int) int {
	components := 0
	seen := make(map[int]bool, n)
	adj := make([][]int, n)
	for _, v := range edges {
		adj[v[0]] = append(adj[v[0]], v[1])
		adj[v[1]] = append(adj[v[1]], v[0])
	}

	var pop int

	for i := range adj {
		if !seen[i] {
			seen[i] = true
			components++
			q := []int{i}
			for len(q) > 0 {
				pop, q = q[0], q[1:]
				for _, v := range adj[pop] {
					if !seen[v] {
						seen[v] = true
						q = append(q, v)
					}
				}
			}
		}
	}

	return components
}

func countComponents2(n int, edges [][]int) int {
	count := 0

	u := NewUnion(n)

	for _, v := range edges {
		if !u.Connected(v[0], v[1]) {
			count++
		}
		u.Union(v[0], v[1])
	}

	return n - count
}

func main() {
	x := 5
	nums := [][]int{{0, 1}, {1, 2}, {3, 4}}
	fmt.Println(countComponents(x, nums))
}
