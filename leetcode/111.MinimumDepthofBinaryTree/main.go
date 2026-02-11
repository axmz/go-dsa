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
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	leftDepth := minDepth(root.Left)
	if leftDepth == 0 {
		return 1 + minDepth(root.Right)
	}
	rightDepth := minDepth(root.Right)
	if rightDepth == 0 {
		return 1 + leftDepth
	}
	return 1 + min(leftDepth, rightDepth)
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
