package main

import (
	"fmt"
	. "godsa/faang/utils"
)

func dfs(root *TreeNode, ch chan int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		ch <- root.Val
		return
	} else {
		dfs(root.Left, ch)
		dfs(root.Right, ch)
	}
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	res := true

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		dfs(root1, ch1)
		close(ch1)
	}()
	go func() {
		dfs(root2, ch2)
		close(ch2)
	}()

	for {
		leaf1, ok1 := <-ch1
		fmt.Println("root 1", ok1, leaf1)
		leaf2, ok2 := <-ch2
		fmt.Println("root 2", ok2, leaf2)

		if ok1 != ok2 || leaf1 != leaf2 {
			res = false
			break
		}

		if !ok1 || !ok2 {
			break
		}
	}

	return res
}

func main() {
	root1 := []any{3, 5, 1, 6, 2, 9, 8, nil, nil, 7, 4}
	root2 := []any{3, 5, 1, 6, 7, 4, 2, nil, nil, nil, nil, nil, nil, 9, 8}

	// There is a better algo, Morris traversal. It doesn't use stack. Space: O(1)
	res := leafSimilar(CreateTree(root1), CreateTree(root2))
	fmt.Println(res)
}
