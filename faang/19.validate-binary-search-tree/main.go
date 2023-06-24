package main

import (
	"fmt"
	. "godsa/faang/utils"
	"math"
)

func isValidBST(root *TreeNode) bool {

	return checkConstraints(root, math.MaxInt, math.MinInt)
}

func checkConstraints(node *TreeNode, st, gt int) bool {
	if node == nil {
		return true
	}

	if node.Val >= st || node.Val <= gt {
		return false
	}

	return checkConstraints(node.Left, node.Val, gt) && checkConstraints(node.Right, st, node.Val)
}

func main() {
	// root := []interface{}{1, 2, 3, 4, 5, 6, 7}
	// root := []interface{}{3, 9, 20, nil, nil, 15, 7}
	// root := []interface{}{5, 1, 4, nil, nil, 3, 6}
	// root := []interface{}{2, 1, 3}
	root := []interface{}{2, 2, 2}
	tree := CreateTree(root)
	res := isValidBST(tree)
	fmt.Println(res)
}
