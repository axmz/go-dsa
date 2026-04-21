package main

import "fmt"

type NumMatrix struct {
	tree [][]int
	nums [][]int
	m, n int
}

func Constructor(matrix [][]int) NumMatrix {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return NumMatrix{}
	}

	m, n := len(matrix), len(matrix[0])

	tree := make([][]int, m+1)
	for i := range tree {
		tree[i] = make([]int, n+1)
	}

	nums := make([][]int, m)
	for i := range nums {
		nums[i] = make([]int, n)
	}

	nm := NumMatrix{tree, nums, m, n}

	// build tree
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			nm.Update(i, j, matrix[i][j])
		}
	}

	return nm
}

func (this *NumMatrix) Update(row int, col int, val int) {
	delta := val - this.nums[row][col]
	this.nums[row][col] = val

	for i := row + 1; i <= this.m; i += i & -i {
		for j := col + 1; j <= this.n; j += j & -j {
			this.tree[i][j] += delta
		}
	}
}

func (this *NumMatrix) sum(row int, col int) int {
	res := 0
	for i := row; i > 0; i -= i & -i {
		for j := col; j > 0; j -= j & -j {
			res += this.tree[i][j]
		}
	}
	return res
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.sum(row2+1, col2+1) -
		this.sum(row1, col2+1) -
		this.sum(row2+1, col1) +
		this.sum(row1, col1)
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * obj.Update(row,col,val);
 * param_2 := obj.SumRegion(row1,col1,row2,col2);
 */

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
