package main

import (
	"fmt"
)

type boxes map[int]map[byte]bool
type board [][]byte

var digits = []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}

func getBoxId(row, col int) (id int) {
	colId := col / 3
	rowId := row / 3 * 3
	return rowId + colId
}

func fillBoxesMap(brd board) (b boxes) {
	b = make(boxes)
	for i, row := range brd {
		for j, col := range row {
			id := getBoxId(i, j)
			// map inside map is nil if not instantiated. this causes problems
			if _, ok := b[id]; !ok {
				b[id] = make(map[byte]bool)
			}
			if col == '.' {
				continue
			}
			if _, ok := b[id]; ok {
				b[id][col] = true
			} else {
				box := map[byte]bool{col: true}
				b[id] = box
			}
		}
	}

	return b
}

func isValid(digit byte, row, col int, b boxes, brd board) bool {
	boxId := getBoxId(row, col)

	if exists := b[boxId][digit]; exists {
		return false
	}

	for _, v := range brd[row] {
		if v == digit {
			return false
		}
	}

	for _, v := range brd {
		if v[col] == digit {
			return false
		}
	}

	return true
}

func backtrack(brd board, row, col int, b boxes) bool {
	if row == len(brd) {
		return true
	}

	if brd[row][col] == '.' {
		for _, digit := range digits {
			if isValid(digit, row, col, b, brd) {
				brd[row][col] = digit
				boxId := getBoxId(row, col)
				b[boxId][digit] = true
				nextRow := row
				nextCol := col
				if col < len(brd[0])-1 {
					nextCol++
				} else {
					nextCol = 0
					nextRow++
				}
				if backtrack(brd, nextRow, nextCol, b) {
					return true
				} else {
					b[boxId][digit] = false
					brd[row][col] = '.'
				}
			}
		}
		return false
	}

	if col < len(brd[0])-1 { // better to col < 9 this is more readable
		return backtrack(brd, row, col+1, b)
	} else {
		return backtrack(brd, row+1, 0, b)
	}
}

func solveSudoku(brd board) {
	b := fillBoxesMap(brd) // better to have boxes as [9]int. easier than map
	// better to create rows and cols memos to not search in the board all the time
	backtrack(brd, 0, 0, b)
}

func print(brd board) {
	fmt.Printf("\n=====================\n")
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			fmt.Printf("%c ", brd[y][x])
			if x%3 == 2 && x < 8 {
				fmt.Printf("| ")
			}
		}
		fmt.Printf("\n")
		if y%3 == 2 && y < 8 {
			fmt.Printf("-----   -----   -----\n")
		}
	}
	fmt.Printf("=====================\n\n")
}
func set(arr []int, idx, num int) {
	res := (1 << (num - 1))
	arr[idx] |= res
}

func main() {
	arr := make([]int, 9)
	set(arr, 6, 6)
	// brd :=
	// 	board{
	// 		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
	// 		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
	// 		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
	// 		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
	// 		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
	// 		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
	// 		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
	// 		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
	// 		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	// 	}

	brd :=
		board{
			{'.', '.', '9', '7', '4', '8', '.', '.', '.'},
			{'7', '.', '.', '.', '.', '.', '.', '.', '.'},
			{'.', '2', '.', '1', '.', '9', '.', '.', '.'},
			{'.', '.', '7', '.', '.', '.', '2', '4', '.'},
			{'.', '6', '4', '.', '1', '.', '5', '9', '.'},
			{'.', '9', '8', '.', '.', '.', '3', '.', '.'},
			{'.', '.', '.', '8', '.', '3', '.', '2', '.'},
			{'.', '.', '.', '.', '.', '.', '.', '.', '6'},
			{'.', '.', '.', '2', '7', '5', '9', '.', '.'},
		}

	solveSudoku(brd)
	print(brd)
}
