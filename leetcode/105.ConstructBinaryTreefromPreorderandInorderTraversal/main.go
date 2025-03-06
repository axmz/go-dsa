package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	m := make(map[int]int)
	for i, v := range inorder {
		m[v] = i
	}

	rootIdx := 0
	var build func(l, r int) *TreeNode
	build = func(l, r int) *TreeNode {
		if l > r {
			return nil
		}
		rootVal := preorder[rootIdx]
		rootIdx++
		inorderIdx := m[rootVal]
		node := &TreeNode{Val: rootVal,
			Left:  build(l, inorderIdx-1),
			Right: build(inorderIdx+1, r),
		}

		return node
	}

	return build(0, len(preorder)-1)
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
