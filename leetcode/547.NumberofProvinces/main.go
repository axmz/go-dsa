package main

import (
	"fmt"
	. "godsa/leetcode/utils/disjointset"
)

// func dfs(isConnected [][]int, visited []bool, i int) {
// 	visited[i] = true
// 	connections := isConnected[i]
// 	for j, c := range connections {
// 		if !visited[j] && c == 1 {
// 			dfs(isConnected, visited, j)
// 		}
// 	}
// }

// func bfs(isConnected [][]int, visited []bool, i int) {
// 	queue := make([][]int, 0)
// 	queue = append(queue, isConnected[i])
// 	visited[i] = true
// 	var pop []int
// 	for len(queue) > 0 {
// 		pop, queue = queue[0], queue[1:]
// 		for j, c := range pop {
// 			if !visited[j] && c == 1 {
// 				visited[j] = true
// 				queue = append(queue, isConnected[j])
// 			}
// 		}
// 	}
// }

// func findCircleNum(isConnected [][]int) int {
// 	regions := 0
// 	n := len(isConnected)
// 	visited := make([]bool, n)

// 	for i := 0; i < n; i++ {
// 		if !visited[i] {
// 			regions++
// 			bfs(isConnected, visited, i)
// 			// dfs(isConnected, visited, i)
// 		}
// 	}

// 	return regions
// }

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	u := NewUnion(n)
	regions := n

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isConnected[i][j] == 1 && !u.Connected(i, j) {
				u.Union(i, j)
				regions--
			}
		}
	}

	return regions
}

func main() {
	// isConnected := [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}
	// isConnected := [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	isConnected := [][]int{{1, 0, 0, 1}, {0, 1, 1, 0}, {0, 1, 1, 1}, {1, 0, 1, 1}}

	res := findCircleNum(isConnected)
	fmt.Println(res)
}
