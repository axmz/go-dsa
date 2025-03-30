package main

import (
	"fmt"
)

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	adj := make(map[int][]int, n)
	indegree := make([]int, n)

	for _, e := range edges {
		v1, v2 := e[0], e[1]
		adj[v1] = append(adj[v1], v2)
		adj[v2] = append(adj[v2], v1)
		indegree[v1]++
		indegree[v2]++
	}

	q := make([]int, 0)
	for i, v := range indegree {
		if v == 1 {
			q = append(q, i)
		}
	}

	for n > 2 {
		size := len(q)
		n -= size

		for i := 0; i < size; i++ {
			leaf := q[i]

			for _, to := range adj[leaf] {
				indegree[to]--

				if indegree[to] == 1 {
					q = append(q, to)
				}
			}
		}

		q = q[size:]
	}

	return q
}

func main() {
	n := 6
	edges := [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}}
	fmt.Println(findMinHeightTrees(n, edges))
	n = 4
	edges = [][]int{{1, 0}, {1, 2}, {1, 3}}
	fmt.Println(findMinHeightTrees(n, edges))
	n = 11
	edges = [][]int{{0, 1}, {0, 2}, {2, 3}, {0, 4}, {2, 5}, {5, 6}, {3, 7}, {6, 8}, {8, 9}, {9, 10}}
	fmt.Println(findMinHeightTrees(n, edges))
}
