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

func CreateTreeInOrder(m []interface{}) *TreeNode {
	if len(m) == 0 {
		return nil
	}

	return buildInOrderTree(m, 0, len(m)-1)
}

func buildInOrderTree(m []interface{}, start, end int) *TreeNode {
	if start > end {
		return nil
	}

	// Find the middle element to be the root
	mid := (start + end) / 2

	// Convert the middle element to a TreeNode
	var root *TreeNode
	if m[mid] != nil {
		root = &TreeNode{Val: m[mid].(int)}
	}

	// Recursively build the left and right subtrees
	if root != nil {
		root.Left = buildInOrderTree(m, start, mid-1)
		root.Right = buildInOrderTree(m, mid+1, end)
	}

	return root
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
