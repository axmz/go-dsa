package main

import "fmt"

func buildTree(inorder []int, postorder []int) *TreeNode {
	m := make(map[int]int)
	for i, v := range inorder {
		m[v] = i
	}

	rootIdx := len(postorder) - 1
	var build func(l, r int) *TreeNode
	build = func(l, r int) *TreeNode {
		if l > r {
			return nil
		}
		rootVal := postorder[rootIdx]
		rootIdx--
		inorderIdx := m[rootVal]
		node := &TreeNode{Val: rootVal,
			Right: build(inorderIdx+1, r), // order matters
			Left:  build(l, inorderIdx-1),
		}

		return node
	}

	return build(0, len(postorder)-1)
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	idx_map := map[int]int{}
	post_idx := len(postorder) - 1
	var helper func(in_left int, in_right int) *TreeNode
	helper = func(in_left int, in_right int) *TreeNode {
		if in_left > in_right {
			return nil
		}
		root_val := postorder[post_idx]
		root := &TreeNode{Val: root_val}
		index := idx_map[root_val]
		post_idx--
		root.Right = helper(index+1, in_right)
		root.Left = helper(in_left, index-1)
		return root
	}
	for i := 0; i < len(inorder); i++ {
		idx_map[inorder[i]] = i
	}
	return helper(0, len(inorder)-1)
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
