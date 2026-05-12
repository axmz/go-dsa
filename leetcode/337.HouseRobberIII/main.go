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
func rob(root *TreeNode) int {
	var dfs func(node *TreeNode) (int, int)
	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}
		leftRob, leftNotRob := dfs(node.Left)
		rightRob, rightNotRob := dfs(node.Right)

		rob := node.Val + leftNotRob + rightNotRob
		notRob := max(leftRob, leftNotRob) + max(rightRob, rightNotRob)
		return rob, notRob
	}

	return max(dfs(root))
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
