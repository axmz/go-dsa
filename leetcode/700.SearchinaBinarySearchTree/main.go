package main

import (
	"fmt"
	. "godsa/faang/utils"
)

func searchBST(root *TreeNode, val int) *TreeNode {
	for root != nil && root.Val != val {
		if root.Val < val {
			root = root.Right
		} else {
			root = root.Left
		}
	}

	return root
}

func searchBSTRecursive(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == val {
		return root
	}

	if root.Val < val {
		return searchBSTRecursive(root.Right, val)
	} else {
		return searchBSTRecursive(root.Left, val)
	}
}

func main() {
	root := []any{4, 2, 7, 1, 3}
	val := 2
	res := searchBST(CreateTree(root), val)
	fmt.Println(res)
}
