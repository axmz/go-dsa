package main

import "fmt"

func isMirror(r1, r2 *TreeNode) bool {
	if r1 == nil && r2 == nil {
		return true
	}

	if r1 == nil || r2 == nil {
		return false
	}

	return r1.Val == r2.Val && isMirror(r1.Right, r2.Left) && isMirror(r1.Left, r2.Right)
}

func isSymmetric2(root *TreeNode) bool {
	return isMirror(root, root)
}

func isSymmetric(root *TreeNode) bool {
	q := []*TreeNode{root, root}
	for len(q) > 0 {
		r1, r2 := q[0], q[1]
		q = q[2:]

		if r1 == nil && r2 == nil {
			continue
		}

		if r1 == nil || r2 == nil {
			return false
		}

		if r1.Val != r2.Val {
			return false
		}

		q = append(q, r1.Left, r2.Right)
		q = append(q, r1.Right, r2.Left)
	}

	return true
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
