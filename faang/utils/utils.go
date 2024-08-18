package utils

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) String() string {
	sl := TreeToSlice(t)
	return fmt.Sprintf("%v", sl)
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

func TreeToSlice(root *TreeNode) []any {
	if root == nil {
		return []any{}
	}

	var result []any
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node != nil {
			result = append(result, node.Val)
			queue = append(queue, node.Left, node.Right)
		} else {
			result = append(result, nil)
		}
	}

	// Remove trailing nils
	for len(result) > 0 && result[len(result)-1] == nil {
		result = result[:len(result)-1]
	}

	return result
}

func CreateTree(m []interface{}) *TreeNode {
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

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(MaxDepth(root.Left), MaxDepth(root.Right)) + 1
}

func GetTreeHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return GetTreeHeight(root.Left) + 1
}
