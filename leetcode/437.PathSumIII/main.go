package main

import (
	"fmt"
	. "godsa/faang/utils"
)

func pathSum(root *TreeNode, targetSum int) int {
	m := make(map[int]int)
	// m[0] = 1
	count := 0

	preorder(root, targetSum, 0, m, &count)
	return count
}

func preorder(n *TreeNode, targetSum int, cumulativeSum int, m map[int]int, count *int) {
	if n == nil {
		return
	}

	cumulativeSum += n.Val
	if cumulativeSum == targetSum {
		*count++
	}

	dif := cumulativeSum - targetSum
	if v, ok := m[dif]; ok {
		*count += v
	}

	m[cumulativeSum]++
	preorder(n.Left, targetSum, cumulativeSum, m, count)
	preorder(n.Right, targetSum, cumulativeSum, m, count)
	m[cumulativeSum]--
}

func main() {
	// root := []any{4, 2, 7, 1, 3}
	// root := []any{3, 3, nil, 4, 2}
	// root := []any{1}
	// root := []any{3, 1, 4, 3, nil, 1, 5}
	root := []any{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, 5, 1}
	targetSum := 22
	// root := []any{10, 5, -3, 3, 2, nil, 11, 3, -2, nil, 1}
	// targetSum := 8
	tree := CreateTree(root)
	res := pathSum(tree, targetSum)
	fmt.Println(res)

}
