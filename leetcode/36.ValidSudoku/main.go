package main

import (
	"fmt"
)

func isValidSudoku(board [][]byte) bool {
	N := 9
	rows := [9]int{}
	cols := [9]int{}
	boxes := [9]int{}
	for r := 0; r < N; r++ {
		for c := 0; c < N; c++ {
			cell := board[r][c]
			if cell == '.' {
				continue
			} else {
				boxId := r/3*3 + c/3
				num := int(cell - '0')
				mask := 1 << (num - 1)
				if rows[r]&mask != 0 || cols[c]&mask != 0 || boxes[boxId]&mask != 0 {
					return false
				} else {
					rows[r] |= mask
					cols[c] |= mask
					boxes[boxId] |= mask
				}
			}
		}
	}

	return true
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
