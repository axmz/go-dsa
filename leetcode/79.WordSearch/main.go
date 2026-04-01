package main

import "fmt"

func exist(board [][]byte, word string) bool {
	r := len(board)
	c := len(board[0])

	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	var dfs func(i, j, k int, seen map[[2]int]bool) bool
	dfs = func(i, j, k int, seen map[[2]int]bool) bool {
		if board[i][j] != word[k] {
			return false
		}

		if k == len(word)-1 {
			return true
		}

		for _, d := range directions {
			newR, newC := i+d[0], j+d[1]
			if newR >= 0 && newR < r && newC >= 0 && newC < c && !seen[[2]int{newR, newC}] {
				seen[[2]int{newR, newC}] = true
				if dfs(newR, newC, k+1, seen) {
					return true
				}
				delete(seen, [2]int{newR, newC})
			}
		}

		return false
	}

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			seen := make(map[[2]int]bool)
			seen[[2]int{i, j}] = true
			if dfs(i, j, 0, seen) {
				return true
			}
		}
	}

	return false
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
