package main

import (
	"fmt"
	. "godsa/utils/disjointset"
)

func findRedundantConnection(edges [][]int) []int {
	n := len(edges) + 1
	u := NewUnion(n)

	for _, v := range edges {
		if u.Connected(v[0], v[1]) {
			return v
		}
		u.Union(v[0], v[1])
	}
	return nil
}

func main() {
	// edges := [][]int{{1, 2}, {1, 3}, {2, 3}}
	edges := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}
	fmt.Println(findRedundantConnection(edges))

}
