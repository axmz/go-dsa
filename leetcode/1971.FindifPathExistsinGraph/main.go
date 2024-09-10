package main

import (
	"fmt"
	. "godsa/utils/disjointset"
)

// func buildGraph(edges [][]int, n int) map[int][]int {
// 	graph := make(map[int][]int)
// 	for _, edge := range edges {
// 		l, r := edge[0], edge[1]
// 		if v, exists := graph[l]; exists {
// 			graph[l] = append(v, r)
// 		} else {
// 			graph[l] = []int{r}
// 		}
// 		if v, exists := graph[r]; exists {
// 			graph[r] = append(v, l)
// 		} else {
// 			graph[r] = []int{l}
// 		}
// 	}
// 	return graph
// }

// func dfs(visited []bool, graph map[int][]int, source, destination int) bool {
// 	if source == destination {
// 		return true
// 	}
// 	ans := false
// 	visited[source] = true
// 	for _, vtx := range graph[source] {
// 		if !visited[vtx] {
// 			if dfs(visited, graph, vtx, destination) {
// 				ans = true
// 			}
// 		}
// 	}
// 	return ans
// }

// func validPath(n int, edges [][]int, source int, destination int) bool {
// 	visited := make([]bool, n)
// 	graph := buildGraph(edges, n)

// 	return dfs(visited, graph, source, destination)
// }

// func bfs(queue []int, visited []bool, graph map[int][]int, source, destination int) bool {
// 	var pop int
// 	fmt.Println(graph)
// 	for len(queue) > 0 {
// 		pop, queue = queue[0], queue[1:]

// 		if pop == destination {
// 			return true
// 		}
// 		if !visited[pop] {
// 			visited[pop] = true
// 			queue = append(queue, graph[pop]...)
// 		}
// 	}

// 	return false
// }

// func validPath(n int, edges [][]int, source int, destination int) bool {
// 	visited := make([]bool, n)
// 	queue := make([]int, n)
// 	queue = append(queue, source)
// 	graph := buildGraph(edges, n)

// 	return bfs(queue, visited, graph, source, destination)
// }

func validPath(n int, edges [][]int, source int, destination int) bool {
	u := NewUnion(n)

	for _, v := range edges {
		u.Union(v[0], v[1])
	}

	return u.Connected(source, destination)
}

func main() {
	n := 3
	edges := [][]int{{0, 1}, {1, 2}, {2, 0}}
	source := 0
	destination := 2

	// n := 6
	// edges := [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}
	// source := 0
	// destination := 5

	// n := 10
	// edges := [][]int{{2, 9}, {7, 8}, {5, 9}, {7, 2}, {3, 8}, {2, 8}, {1, 6}, {3, 0}, {7, 0}, {8, 5}}
	// source := 1
	// destination := 0

	res := validPath(n, edges, source, destination)
	fmt.Println(res)
}
