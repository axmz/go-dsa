package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertToTreeNode(m []interface{}) []*TreeNode {
	r := make([]*TreeNode, len(m))
	for i, v := range m {
		if v != nil {
			t := v.(int)
			r[i] = &TreeNode{Val: t, Left: nil, Right: nil}
		}
	}
	return r
}

func createTree(m []interface{}) *TreeNode {
	nodes := convertToTreeNode(m)
	kids := nodes[1:]

	for i := range nodes {
		if nodes[i] != nil {
			if len(kids) > 0 {
				nodes[i].Left, kids = kids[0], kids[1:]
			}
			if len(kids) > 0 {
				nodes[i].Right, kids = kids[0], kids[1:]
			}
		}
	}

	return nodes[0]
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

// [3,9,20,null,null,15,7]
// [5,4,7,3,null,2,null,-1,null,9]
func main() {
	// m := []interface{}{3, 9, 20, nil, nil, 15, 7}
	m := []interface{}{5, 4, 7, 3, nil, 2, nil, -1, nil, 9}
	// m := []interface{}{nil}
	p := createTree(m)
	d := maxDepth(p)
	fmt.Println(p, d)
}
