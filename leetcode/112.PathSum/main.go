package main

import "fmt"

func traverse(root *TreeNode, sum, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		if root.Val+sum == targetSum {
			return true
		} else {
			return false
		}
	}

	var l, r bool
	if root.Left != nil {
		l = traverse(root.Left, root.Val+sum, targetSum)
	}

	if root.Right != nil {
		r = traverse(root.Right, root.Val+sum, targetSum)
	}

	if l || r {
		return true
	}

	return false
}

func hasPathSum2(root *TreeNode, targetSum int) bool {
	return traverse(root, 0, targetSum)
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	sum := targetSum - root.Val
	if root.Left == nil && root.Right == nil {
		return sum == 0
	}
	return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
