package main

import "fmt"

func updateMatrix(mat [][]int) [][]int {
	r, c := len(mat), len(mat[0])
	visited := make([][]bool, r)
	for i := range visited {
		visited[i] = make([]bool, c)
	}
	q := make([][2]int, 0)
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if mat[i][j] == 0 {
				q = append(q, [2]int{i, j})
				visited[i][j] = true
			}
		}
	}

	directions := [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	for len(q) > 0 {
		pop := q[0]
		q = q[1:]
		curR, curC := pop[0], pop[1]
		for _, dir := range directions {
			newR, newC := curR+dir[0], curC+dir[1]
			if newR >= 0 && newR < r && newC >= 0 && newC < c && !visited[newR][newC] {
				visited[newR][newC] = true
				mat[newR][newC] = mat[curR][curC] + 1
				q = append(q, [2]int{newR, newC})
			}
		}
	}

	return mat
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
