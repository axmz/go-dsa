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
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0
	var depth func(node *TreeNode) int
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		l := depth(node.Left)
		r := depth(node.Right)
		d := l + r + 1
		maxDiameter = max(d, maxDiameter)
		return max(l, r) + 1
	}
	depth(root)
	return maxDiameter - 1
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
