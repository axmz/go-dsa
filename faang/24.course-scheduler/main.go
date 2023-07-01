package main

import (
	"fmt"
)

func canFinish(numCourses int, prerequisites [][]int) bool {
	adjList := make(map[int][]int)
	inDegree := make([]int, numCourses)
	queue := []int{}
	count := 0

	for _, pair := range prerequisites {
		src := pair[1]
		dst := pair[0]
		adjList[src] = append(adjList[src], dst)
		inDegree[dst] = inDegree[dst] + 1
	}

	for i, v := range inDegree {
		if v == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		count++
		node := queue[0]
		queue = queue[1:]
		dsts := adjList[node]
		for _, v := range dsts {
			inDegree[v] = inDegree[v] - 1
			if inDegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	fmt.Println(adjList, inDegree)
	return count == numCourses
}

func main() {
	numCourses := 6
	prerequisites := [][]int{{1, 0}, {2, 1}, {2, 5}, {0, 3}, {4, 3}, {3, 5}, {4, 5}}
	// prerequisites := [][]int{{0,1}, {1,0}}
	res := canFinish(numCourses, prerequisites)
	fmt.Println(res)
}
