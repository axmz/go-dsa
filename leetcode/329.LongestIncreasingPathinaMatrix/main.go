package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func longestIncreasingPath(matrix [][]int) int {
	h := len(matrix)
	w := len(matrix[0])

	directions := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	memo := make([][]int, h)
	for i := range memo {
		memo[i] = make([]int, w)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(r, c int) int
	dfs = func(r, c int) int {
		if memo[r][c] != -1 {
			return memo[r][c]
		}

		memo[r][c] = 1
		for _, d := range directions {
			nr := r + d[0]
			nc := c + d[1]
			if nr >= 0 && nr < h && nc >= 0 && nc < w && matrix[r][c] < matrix[nr][nc] {
				memo[r][c] = max(memo[r][c], dfs(nr, nc)+1)
			}
		}

		return memo[r][c]
	}

	longest := 0
	for r := 0; r < h; r++ {
		for c := 0; c < w; c++ {
			longest = max(longest, dfs(r, c))
		}
	}

	return longest
}

func longestIncreasingPath3(matrix [][]int) int {
	h := len(matrix)
	w := len(matrix[0])

	directions := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	adj := make(map[int][]int)

	for r := 0; r < h; r++ {
		for c := 0; c < w; c++ {
			for _, d := range directions {
				nr := r + d[0]
				nc := c + d[1]
				if nr >= 0 && nr < h && nc >= 0 && nc < w && matrix[r][c] < matrix[nr][nc] {
					idx := r*w + c
					nidx := nr*w + nc
					adj[idx] = append(adj[idx], nidx)
				}
			}
		}
	}

	memo := make(map[int]int)

	cacheHit := 0
	cacheMis := 0
	var dfs func(node int) int
	dfs = func(node int) int {
		if v, exist := memo[node]; exist {
			cacheHit += 1
			fmt.Println("cache hit")
			return v
		}
		cacheMis += 1
		fmt.Println("cache mis")

		if neighbours, exist := adj[node]; !exist {
			memo[node] = 1
		} else {
			curMax := 0
			for _, n := range neighbours {
				curMax = max(curMax, dfs(n))
			}
			memo[node] = curMax + 1
		}
		return memo[node]
	}

	longest := 1
	for k := range adj {
		longest = max(longest, dfs(k))
	}

	fmt.Println(cacheHit, cacheMis)
	return longest
}

func main() {
	// matrix := [][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}}
	matrix := [][]int{
		{7, 7, 5},
		{2, 4, 6},
		{8, 2, 0}}
	fmt.Println(longestIncreasingPath(matrix))
}
