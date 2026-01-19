package main

import "fmt"

func lowestCommonAncestorIter(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		if p.Val < root.Val && q.Val < root.Val {
			root = root.Left
		} else if p.Val > root.Val && q.Val > root.Val {
			root = root.Right
		} else {
			return root
		}
	}

	return nil
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// this isn't even necessary for BST
	// if root == nil || root == p || root == q {
	// 	return root
	// }

	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}

	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}

	return root
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
