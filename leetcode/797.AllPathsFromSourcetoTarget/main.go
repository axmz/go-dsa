package main

import (
	"fmt"
)

// DFS
// func dfs(graph [][]int, node int, path []int, paths *[][]int) {
// 	path = append(path, node)

// 	if node == len(graph)-1 {
// 		*paths = append(*paths, append([]int(nil), path...))
// 		return
// 	}
// 	for _, next := range graph[node] {
// 		dfs(graph, next, path, paths)
// 	}
// }

// func allPathsSourceTarget(graph [][]int) [][]int {
// 	var paths [][]int
// 	dfs(graph, 0, []int{}, &paths)
// 	return paths
// }

// DFS with memo
// func dfs(graph [][]int, node int, memo map[int][][]int) [][]int {
// 	if m, exists := memo[node]; exists {
// 		return m
// 	}

// 	paths := [][]int{}
// 	if node == len(graph)-1 {
// 		return [][]int{{node}}
// 	}

// 	// not the edgecase
// 	// if len(graph[node]) == 0 {
// 	// 	return [][]int{}
// 	// }

// 	for _, next := range graph[node] {
// 		for _, path := range dfs(graph, next, memo) {
// 			// if len(path) == 0 {
// 			// 	continue
// 			// }

// 			path = append([]int{node}, path...)
// 			paths = append(paths, append([]int(nil), path...))
// 		}
// 	}

// 	memo[node] = paths

// 	return paths
// }

// func allPathsSourceTarget(graph [][]int) [][]int {
// 	memo := make(map[int][][]int)
// 	return dfs(graph, 0, memo)
// }

// BFS
func bfs(graph [][]int, queue [][]int, paths [][]int) [][]int {
	for len(queue) > 0 {
		var path []int
		path, queue = queue[0], queue[1:]
		node := path[len(path)-1]

		for _, next := range graph[node] {
			newPath := append([]int{}, path...) // Ensure a new copy of the path is made
			newPath = append(newPath, next)
			if next == len(graph)-1 {
				paths = append(paths, newPath)
			} else {
				queue = append(queue, newPath)
			}
		}
	}

	return paths
}

func allPathsSourceTarget(graph [][]int) [][]int {
	queue := make([][]int, 0)
	queue = append(queue, []int{0})
	var paths [][]int
	return bfs(graph, queue, paths)
}

func main() {
	// graph := [][]int{{1, 2}, {3}, {3}, {}}
	graph := [][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}}
	res := allPathsSourceTarget(graph)
	fmt.Println(res)
}
