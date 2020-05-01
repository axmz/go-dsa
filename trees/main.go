package main

import (
	"fmt"
	"godsa/trees/bst"
)

func printPreOrder(n *bst.Node) {
	if n == nil {
		return
	}
	fmt.Printf("%d ", n.Value)
	printPreOrder(n.Left)
	printPreOrder(n.Right)
}

func main() {
	var tree = bst.NewBst()
	tree.Insert(9)
	tree.Insert(4)
	tree.Insert(6)
	tree.Insert(20)
	tree.Insert(170)
	tree.Insert(15)
	tree.Insert(1)

	// tree.Remove(1)
	// tree.Remove(4)
	// tree.Remove(9)
	// tree.Remove(4)
	// tree.Remove(4)

	// n, p, e := tree.Lookup(4)
	// n.Print()
	// p.Print()
	// println(e)

	// tree.Print()
	// printPreOrder(tree.Root)

	// BFS
	// bfs := tree.BFS()
	// fmt.Println(bfs)

	// BFSR
	// bfsr := tree.BFSR(nil, []int{})
	// fmt.Println(bfsr)

	// DFS
	inOrder := tree.InOrder(tree.Root, &[]int{})
	preOrder := tree.PreOrder(tree.Root, &[]int{})
	postOrder := tree.PostOrder(tree.Root, &[]int{})
	fmt.Println(inOrder)
	fmt.Println(preOrder)
	fmt.Println(postOrder)
}
