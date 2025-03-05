package main

import (
	"fmt"
	. "godsa/utils/tree"
)

func countUnivalSubtrees(root *TreeNode) int {
	var dfs func(root *TreeNode) bool
	sum := 0
	dfs = func(root *TreeNode) bool {
		if root == nil {
			return true
		}

		l := dfs(root.Left)
		r := dfs(root.Right)
		if l && r {
			if root.Left != nil && root.Left.Val != root.Val || root.Right != nil && root.Right.Val != root.Val {
				return false
			}
			sum++
			return true
		}

		return false

	}

	dfs(root)
	return sum

}

func main() {
	root := []interface{}{5, 1, 3, 1, 1, 1}
	// root := []interface{}{5, 1, 5, 5, 5, nil, 5}
	tree := CreateTree(root)
	fmt.Println(countUnivalSubtrees(tree))
}
