package main

import (
	"fmt"
	. "godsa/utils/tree"
	"math"
)

func countNodes(root *TreeNode) int {
	h := GetTreeHeight(root)
	nodesOnLastLevel := findNodesOnLastLevel(root, h) + 1
	nodesExceptLastLevel := int(math.Pow(2, float64(h-1))) - 1

	return nodesExceptLastLevel + nodesOnLastLevel
}

func findNodesOnLastLevel(root *TreeNode, height int) int {
	l := 0
	r := int(math.Pow(2, float64(height))/2) - 1

	for l < r {
		idxToSearch := int(math.Ceil((float64(l) + float64(r)) / 2))

		if nodeExists(root, idxToSearch, height) != nil {
			l = idxToSearch
		} else {
			r = idxToSearch - 1
		}
	}

	return l
}

func nodeExists(root *TreeNode, idxToSearch, height int) *TreeNode {
	l := 0
	r := int(math.Pow(2, float64(height))/2) - 1
	curLevel := 1
	node := root
	for curLevel < height {
		m := int(math.Ceil((float64(l) + float64(r)) / 2))
		if idxToSearch >= m {
			node = node.Right
			l = m
		} else {
			node = node.Left
			r = m - 1
		}
		curLevel++
	}

	return node
}

func main() {
	root := []interface{}{1, 2, 3, 4, 5, 6, 7}
	// root := []interface{}{3, 9, 20, nil, nil, 15, 7}
	tree := CreateTree(root)
	res := countNodes(tree)
	fmt.Println(res)
}
