package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b, c int) int {
	if a > b && a > c {
		return a
	}

	if b > c {
		return b
	}

	return c
}

func maxPathSum(root *TreeNode) int {
	maxVal := -10000

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		v := node.Val
		l := dfs(node.Left)
		lv := l + v
		r := dfs(node.Right)
		rv := r + v

		p := max(lv, rv, v)
		maxVal = max(maxVal, p, l+r+v)

		return p
	}

	dfs(root)

	return maxVal
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
