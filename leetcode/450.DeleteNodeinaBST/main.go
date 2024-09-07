package main

import (
	"fmt"
	. "godsa/utils/tree"
)

func isLeaf(root *TreeNode) bool {
	return root.Left == nil && root.Right == nil
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	switch {
	case root == nil:
		return root
	case key < root.Val:
		root.Left = deleteNode(root.Left, key)
	case key > root.Val:
		root.Right = deleteNode(root.Right, key)
	default:
		switch {
		case isLeaf(root):
			return nil
		case root.Right == nil:
			return root.Left
		case root.Left == nil:
			return root.Right
		default:
			successor := root.Right
			for successor.Left != nil {
				successor = successor.Left
			}
			root.Val = successor.Val
			root.Right = deleteNode(root.Right, successor.Val)
		}
	}

	return root
}

func main() {
	// root := []any{1, 2, 3, 4, 5, 6, 7}
	// root := []any{5, 3, 6, 2, 4, nil, 7}
	// root := []any{5, 3, 6, 2, 4, nil, 7}
	// root := []any{0}
	root := []any{5, 3, 6, 2, 4, nil, 7}
	// root := []any{50, 30, 70, nil, 40, 60, 80}
	tree := CreateTree(root)
	key := 7
	res := deleteNode(tree, key)
	fmt.Println(res)
}
