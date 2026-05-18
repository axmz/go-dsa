package main

import "fmt"

func buildMatrix(k int, rowConditions [][]int, colConditions [][]int) [][]int {

	topologicalSort := func(conditions [][]int) []int {
		inDegree := make([]int, k+1)
		graph := make(map[int][]int)

		for _, c := range conditions {
			u, v := c[0], c[1]
			graph[u] = append(graph[u], v)
			inDegree[v]++
		}

		queue := make([]int, 0)
		for i := 1; i <= k; i++ {
			if inDegree[i] == 0 {
				queue = append(queue, i)
			}
		}

		order := make([]int, 0)
		for len(queue) > 0 {
			u := queue[0]
			queue = queue[1:]
			order = append(order, u)

			for _, v := range graph[u] {
				inDegree[v]--
				if inDegree[v] == 0 {
					queue = append(queue, v)
				}
			}
		}

		if len(order) != k {
			return nil // cycle detected
		}
		return order
	}

	rowOrder := topologicalSort(rowConditions)
	colOrder := topologicalSort(colConditions)

	if rowOrder == nil || colOrder == nil {
		return [][]int{}
	}

	matrix := make([][]int, k)
	for i := range matrix {
		matrix[i] = make([]int, k)
	}

	rowPos := make([]int, k+1)
	colPos := make([]int, k+1)

	for i, num := range rowOrder {
		rowPos[num] = i
	}

	for j, num := range colOrder {
		colPos[num] = j
	}

	for num := 1; num <= k; num++ {
		r := rowPos[num]
		c := colPos[num]
		matrix[r][c] = num
	}

	return matrix
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
