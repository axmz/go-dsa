package main

import (
	"fmt"
)

// In this problem we eliminate 2^n time and m*n state.
// We represent state as two rows (bitmask)
// So, instead of going on seat-by-seat we precompute entire row configurations
func maxStudents(seats [][]byte) int {
	rows, cols := len(seats), len(seats[0])

	// convert seats to bitmask representation
	seatsMask := make([]int, rows)
	for i := 0; i < rows; i++ {
		mask := 0
		for j := 0; j < cols; j++ {
			if seats[i][j] == '#' {
				mask |= 1 << j
			}
		}
		seatsMask[i] = mask
	}

	// precompute valid row states
	valid := make([][]int, rows)
	for i := 0; i < rows; i++ {
		for mask := 0; mask < (1 << cols); mask++ {
			// check if can be seated && no adjacent students
			if mask&seatsMask[i] == 0 && mask&(mask>>1) == 0 {
				valid[i] = append(valid[i], mask)
			}
		}
	}

	memo := make(map[[2]int]int)

	var dfs func(rowIdx, prevMask int) int
	dfs = func(rowIdx, prevMask int) int {
		if rowIdx == rows {
			return 0
		}

		key := [2]int{rowIdx, prevMask}
		if v, ok := memo[key]; ok {
			return v
		}

		best := 0

		for _, cur := range valid[rowIdx] {
			if !compatible(prevMask, cur) {
				continue
			}

			// each 1 represent a student
			students := countOnes(cur)
			total := students + dfs(rowIdx+1, cur)

			if total > best {
				best = total
			}
		}

		memo[key] = best
		return best
	}

	return dfs(0, 0)
}

func compatible(prev, cur int) bool {
	// Prevent cheating with upper-left and upper-right neighbors from the previous row.
	return (prev<<1)&cur == 0 && (prev>>1)&cur == 0
}

func countOnes(x int) int {
	c := 0
	for x > 0 {
		x &= x - 1
		c++
	}
	return c
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
