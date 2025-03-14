package main

import (
	"fmt"
	. "godsa/utils/disjointset"
)

func validTree(n int, edges [][]int) bool {
	if len(edges) != n-1 {
		return false
	}

	seen := make(map[int]bool)
	adj := make([][]int, n)

	for _, v := range edges {
		adj[v[0]] = append(adj[v[0]], v[1])
		adj[v[1]] = append(adj[v[1]], v[0])
	}

	var dfs func(e int)
	dfs = func(e int) {
		if seen[e] {
			return
		}
		seen[e] = true
		for _, v := range adj[e] {
			dfs(v)
		}
	}

	dfs(0)
	return len(seen) == n
}

func validTree2(n int, edges [][]int) bool {
	u := NewUnion(n)
	if len(edges) != n-1 {
		return false
	}

	for _, e := range edges {
		if u.Connected(e[0], e[1]) {
			return false
		} else {
			u.Union(e[0], e[1])
		}
	}

	return true
}

func main() {
	// n := 5
	// edges := [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4}}
	// n := 4
	// edges := [][]int{{0, 1}, {2, 3}, {1, 2}}
	// n := 5
	// edges := [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}}
	n := 5
	edges := [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}}
	// n := 5
	// edges := [][]int{{0, 1}, {0, 4}, {1, 4}, {2, 3}}
	fmt.Println(validTree(n, edges))
}
