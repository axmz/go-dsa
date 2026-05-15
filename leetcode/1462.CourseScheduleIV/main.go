package main

import "fmt"

func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	adj := make([][]int, numCourses)
	indegree := make([]int, numCourses)
	for _, p := range prerequisites {
		prereq := p[0]
		crs := p[1]
		adj[prereq] = append(adj[prereq], crs)
		indegree[crs]++
	}

	connections := make([][]bool, numCourses)
	for i := range connections {
		connections[i] = make([]bool, numCourses)
	}

	q := make([]int, 0)
	for i, deg := range indegree {
		if deg == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		prereq := q[0]
		q = q[1:]
		for _, crs := range adj[prereq] {
			connections[prereq][crs] = true
			for i := 0; i < numCourses; i++ {
				if connections[i][prereq] {
					connections[i][crs] = true
				}
			}
			indegree[crs]--
			if indegree[crs] == 0 {
				q = append(q, crs)
			}
		}
	}

	res := make([]bool, len(queries))
	for i, q := range queries {
		res[i] = connections[q[0]][q[1]]
	}
	return res
}

func checkIfPrerequisite2(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	adj := make([][]int, numCourses)
	for _, p := range prerequisites {
		adj[p[0]] = append(adj[p[0]], p[1])
	}

	prereqs := make([][]bool, numCourses)
	for i := range prereqs {
		prereqs[i] = make([]bool, numCourses)
	}

	var dfs func(src, node int)
	dfs = func(src, node int) {
		for _, next := range adj[node] {
			if !prereqs[src][next] {
				prereqs[src][next] = true
				dfs(src, next)
			}
		}
	}

	for i := 0; i < numCourses; i++ {
		dfs(i, i)
	}

	res := make([]bool, len(queries))
	for i, q := range queries {
		res[i] = prereqs[q[0]][q[1]]
	}

	return res
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
