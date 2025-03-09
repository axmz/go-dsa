package main

import "fmt"

func leadsToDestination(n int, edges [][]int, source int, destination int) bool {
	visited := make(map[int]bool, n)
	adj := make(map[int][]int, n)
	for _, v := range edges {
		adj[v[0]] = append(adj[v[0]], v[1])
	}

	var dfs func(source int) bool
	dfs = func(source int) bool {
		// important line
		if len(adj[source]) == 0 {
			return source == destination
		}
		if v, exists := visited[source]; exists {
			return v
		}
		visited[source] = false
		for _, v := range adj[source] {
			if !dfs(v) {
				return false
			}
		}
		visited[source] = true
		return visited[source]
	}

	return dfs(source)
}

func main() {
	n := 3
	edges := [][]int{{0, 1}, {0, 2}}
	source := 0
	destination := 2
	// n := 4
	// edges := [][]int{{0, 1}, {0, 2}, {1, 3}, {2, 3}}
	// source := 0
	// destination := 3

	fmt.Println(leadsToDestination(n, edges, source, destination))
}
