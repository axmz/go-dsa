package main

import "fmt"

// // Time Limit Exceeded
// func minReorder(n int, connections [][]int) int {
// 	var changes, l, r int
// 	var pop []int
// 	leadsToZero := make(map[int]bool, n)
// 	leadsToZero[0] = true

// 	for len(connections) > 0 {
// 		pop, connections = connections[0], connections[1:]
// 		l, r = pop[0], pop[1]
// 		if leadsToZero[r] {
// 			leadsToZero[l] = true
// 		} else if leadsToZero[l] {
// 			leadsToZero[r] = true
// 			changes++
// 		} else {
// 			connections = append(connections, pop)
// 		}
// 	}

// 	return changes
// }

func createAdjList(connections [][]int) map[int][][]int {
	adj := map[int][][]int{}
	var l, r int
	var isOriginal, isNotOriginal int = 1, -1
	for _, con := range connections {
		l, r = con[0], con[1]
		adj[l] = append(adj[l], []int{r, isOriginal})
		adj[r] = append(adj[r], []int{l, isNotOriginal})
	}

	return adj
}

// BFS
// func bfs(adj map[int][][]int, queue []int, visited []bool, count *int) {
// 	var pop int
// 	var conns [][]int
// 	var direction, isOriginal int
// 	for len(queue) > 0 {
// 		pop, queue = queue[0], queue[1:]
// 		conns = adj[pop]
// 		visited[pop] = true

// 		for _, con := range conns {
// 			direction = con[0]
// 			isOriginal = con[1]
// 			if !visited[direction] {
// 				if isOriginal == 1 {
// 					*count++
// 				}
// 				queue = append(queue, direction)
// 			}
// 		}
// 	}
// }

// func minReorder(n int, connections [][]int) int {
// 	adj := createAdjList(connections)
// 	queue := []int{0}
// 	count := 0
// 	visited := make([]bool, n)

// 	fmt.Println(adj)
// 	bfs(adj, queue, visited, &count)

// 	return count
// }

// DFS
func dfs(node int, adj map[int][][]int, count *int, parent int) {
	var conns [][]int
	var direction, isOriginal int
	conns = adj[node]

	for _, con := range conns {
		direction = con[0]
		isOriginal = con[1]
		// if direction = parent it means it is a visited node
		// no need for visited array because when n node and n-1 edges it means there can be no cycles
		if parent != direction {
			if isOriginal == 1 {
				*count++
			}

			dfs(direction, adj, count, node)
		}
	}
}

func minReorder(n int, connections [][]int) int {
	adj := createAdjList(connections)
	count := 0

	fmt.Println(adj)
	dfs(0, adj, &count, -1)

	return count
}

func main() {
	n := 6
	connections := [][]int{{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5}}
	// n := 5
	// connections := [][]int{{1, 0}, {1, 2}, {3, 2}, {3, 4}}
	// n := 3
	// connections := [][]int{{1, 0}, {2, 0}}
	// n := 3
	// connections := [][]int{{0, 1}, {0, 2}}

	res := minReorder(n, connections)
	fmt.Println(res)

}
