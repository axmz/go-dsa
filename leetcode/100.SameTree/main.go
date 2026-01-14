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
func isSameTree(p *TreeNode, q *TreeNode) bool {
	ps := []*TreeNode{p}
	qs := []*TreeNode{q}

	for len(ps) > 0 && len(qs) > 0 {
		pNode := ps[0]
		qNode := qs[0]
		ps = ps[1:]
		qs = qs[1:]
		if pNode == nil && qNode == nil {
			continue
		}
		if pNode == nil || qNode == nil {
			return false
		}
		if pNode.Val != qNode.Val {
			return false
		}
		ps = append(ps, pNode.Left)
		ps = append(ps, pNode.Right)
		qs = append(qs, qNode.Left)
		qs = append(qs, qNode.Right)
	}

	return len(ps) == 0 && len(qs) == 0
}

func isSameTreeRec(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
