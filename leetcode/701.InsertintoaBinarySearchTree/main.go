package main

import "fmt"

func insertIntoBSTIterativePtr(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	current := &root
	for *current != nil {
		if val < (*current).Val {
			current = &(*current).Left
		} else {
			current = &(*current).Right
		}
	}
	*current = &TreeNode{Val: val}

	return root
}

func insertIntoBSTIterative(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	current := root
	for {
		if val < current.Val {
			if current.Left == nil {
				current.Left = &TreeNode{Val: val}
				break
			}
			current = current.Left
		} else {
			if current.Right == nil {
				current.Right = &TreeNode{Val: val}
				break
			}
			current = current.Right
		}
	}

	return root
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	if val < root.Val {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		root.Right = insertIntoBST(root.Right, val)
	}

	return root
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
