package main

import (
	"fmt"
	. "godsa/utils/graph"
)

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

// DFS
// func dfs(node *Node, nodeAdjList map[int]*Node) *Node {
// 	if n, exists := nodeAdjList[node.Val]; exists {
// 		return n
// 	}

// 	nodeCopy := &Node{
// 		Val:       node.Val,
// 		Neighbors: []*Node{},
// 	}

// 	nodeAdjList[node.Val] = nodeCopy

// 	for _, ne := range node.Neighbors {
// 		newNeighbour := dfs(ne, nodeAdjList)
// 		nodeCopy.Neighbors = append(nodeCopy.Neighbors, newNeighbour)
// 	}

// 	return nodeCopy
// }

// func cloneGraph(node *Node) *Node {
// 	if node == nil {
// 		return node
// 	}
// 	nodeAdjList := map[int]*Node{}
// 	return dfs(node, nodeAdjList)
// }

// BFS
func bfs(queue []*Node, clones map[*Node]*Node) {
	var pop *Node
	var cur *Node

	for len(queue) > 0 {
		pop, queue = queue[0], queue[1:]

		if n, exists := clones[pop]; exists {
			cur = n
		} else {
			cur = &Node{Val: pop.Val, Neighbors: []*Node{}}
			clones[pop] = cur
		}

		for _, neighbour := range pop.Neighbors {
			if _, exists := clones[neighbour]; !exists {
				clones[neighbour] = &Node{Val: neighbour.Val, Neighbors: []*Node{}}
				queue = append(queue, neighbour)
			}
			cur.Neighbors = append(cur.Neighbors, clones[neighbour])
		}
	}
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return node
	}
	clones := map[*Node]*Node{}
	queue := []*Node{node}
	bfs(queue, clones)

	return clones[node]
}

func main() {
	adjList := [][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}}
	nodeAdjList := map[int]*Node{}
	graph := BuildGraph(1, adjList, nodeAdjList)

	res := cloneGraph(graph)
	fmt.Println(res)
}
