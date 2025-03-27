package main

import "fmt"

func findOrder(numCourses int, prerequisites [][]int) []int {
	res := []int{}
	adj := make(map[int][]int, numCourses)
	indegree := make(map[int]int)
	q := make([]int, 0)

	for _, v := range prerequisites {
		indegree[v[1]]++
		adj[v[0]] = append(adj[v[0]], v[1])
	}

	for i := 0; i < numCourses; i++ {
		if indegree[i] == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		pop := q[0]
		q = q[1:]
		res = append([]int{pop}, res...)
		for _, v := range adj[pop] {
			indegree[v]--
			if indegree[v] == 0 {
				q = append(q, v)
			}
		}
	}

	if len(res) == numCourses {
		return res
	}

	return []int{}
}

func main() {
	n := 4
	nums := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}
	fmt.Println(findOrder(n, nums))
}
