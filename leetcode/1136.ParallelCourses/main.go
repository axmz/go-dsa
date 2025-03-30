package main

import "fmt"

func minimumSemesters2(n int, relations [][]int) int {
	adj := make(map[int][]int, n+1)
	indegree := make([]int, n+1)
	q := make([]int, 0)

	semestersCount := 0
	sortedCount := 0

	for _, v := range relations {
		c1 := v[0]
		c2 := v[1]
		adj[c1] = append(adj[c1], c2)
		indegree[c2]++
	}

	for i, v := range indegree {
		if i == 0 {
			continue
		}

		if v == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		size := len(q)
		semestersCount++

		newQ := []int{}
		for i := 0; i < size; i++ {
			pop := q[0]
			q = q[1:]
			sortedCount++
			for _, node := range adj[pop] {
				indegree[node]--
				if indegree[node] == 0 {
					newQ = append(newQ, node)
				}
			}
		}
		q = newQ
	}

	if sortedCount == n {
		return semestersCount
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func minimumSemesters3(n int, relations [][]int) int {

	adj := make(map[int][]int, n+1)
	for _, v := range relations {
		adj[v[0]] = append(adj[v[0]], v[1])
	}

	visitedLoop := make(map[int]bool, n+1)

	var dfsLoopCheck func(node int) bool
	dfsLoopCheck = func(node int) bool {
		if v, exist := visitedLoop[node]; exist {
			return v
		}
		visitedLoop[node] = true
		for _, v := range adj[node] {
			if dfsLoopCheck(v) {
				return true
			}
		}
		visitedLoop[node] = false

		return visitedLoop[node]
	}

	visitedLongestPath := make([]int, n+1)

	var longestPath func(node int) int
	longestPath = func(node int) int {
		if visitedLongestPath[node] != 0 {
			return visitedLongestPath[node]
		}

		visitedLongestPath[node] = 1
		for _, v := range adj[node] {
			visitedLongestPath[node] = max(longestPath(v)+1, visitedLongestPath[node])
		}

		return visitedLongestPath[node]
	}

	for k := range adj {
		if dfsLoopCheck(k) {
			return -1
		}
	}

	l := 0
	for k := range adj {
		l = max(longestPath(k), l)
	}

	return l
}

func minimumSemesters(n int, relations [][]int) int {
	adj := make(map[int][]int, n+1)
	for _, v := range relations {
		adj[v[0]] = append(adj[v[0]], v[1])
	}

	visited := make([]int, n+1)

	var dfs func(node int) int
	dfs = func(node int) int {
		if visited[node] != 0 {
			return visited[node]
		}

		visited[node] = -1
		mx := 1
		for _, n := range adj[node] {
			r := dfs(n)
			if r == -1 {
				return -1
			}
			mx = max(r+1, mx)
		}
		visited[node] = mx

		return visited[node]
	}

	mx := 0
	for k := range adj {
		res := dfs(k)
		if res == -1 {
			return -1
		}

		mx = max(res, mx)
	}

	return mx
}

func main() {
	n := 3
	relations := [][]int{{1, 3}, {2, 3}}
	fmt.Println(minimumSemesters(n, relations))

	n = 3
	relations = [][]int{{1, 3}, {2, 3}, {3, 1}}
	fmt.Println(minimumSemesters(n, relations))
}
