package main

import "fmt"

func kWeakestRows(mat [][]int, k int) []int {
	res := make([]int, 0, k)
	rows := len(mat)
	cols := len(mat[0])

	for c := 0; c < cols; c++ {
		for r := 0; r < rows; r++ {
			if mat[r][c] == 0 && (c == 0 || mat[r][c-1] != 0) {
				res = append(res, r)
				if len(res) == k {
					return res
				}
			}
		}
	}

	for r := 0; r < rows; r++ {
		if mat[r][cols-1] == 0 {
			continue
		}
		res = append(res, r)
		if len(res) == k {
			return res
		}
	}

	return res
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
