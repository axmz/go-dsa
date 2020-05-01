package bst

import (
	"encoding/json"
	"fmt"
)

// Node structure of the node
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// Insert inserts a num into BST
func (n *Node) Insert(num int) {
	if n.Value <= num {
		if n.Right == nil {
			n.Right = &Node{Value: num}
		} else {
			n.Right.Insert(num)
		}
	} else {
		if n.Left == nil {
			n.Left = &Node{Value: num}
		} else {
			n.Left.Insert(num)
		}
	}
}

// Lookup search for a value in BST
func (n *Node) Lookup(num int, parent *Node) (node *Node, parentNode *Node, err error) {
	if n.Value == num {
		return n, parent, err
	} else if n.Value <= num {
		if n.Right == nil {
			return node, parentNode, fmt.Errorf("Node not found")
		}
		return n.Right.Lookup(num, n)
	} else {
		if n.Left == nil {
			return node, parentNode, fmt.Errorf("Node not found")
		}
		return n.Left.Lookup(num, n)
	}
}

// Remove removes the node from the tree
func (n Node) Remove(num int) (node Node, parentNode *Node, err error) {

	if node.Right == nil && node.Left == nil { // no right, no left - remove entirely
		if parentNode.Right.Value == node.Value {
			parentNode.Right = nil
		}
		if parentNode.Left.Value == node.Value {
			parentNode.Left = nil
		}
	} else if node.Right == nil { // no right - connect parent with left
		if parentNode.Right.Value == node.Value {
			parentNode.Right = node.Left
		}
		if parentNode.Left.Value == node.Value {
			parentNode.Left = node.Left
		}
	} else if node.Right != nil { // right - connect parent with leftmost
		leftmost := node.Right
		for {
			if leftmost.Left != nil {
				leftmost = leftmost.Left
			} else {
				break
			}
		}
		if parentNode.Right.Value == node.Value {
			parentNode.Right = leftmost
		}
		if parentNode.Left.Value == node.Value {
			parentNode.Left = leftmost
		}
	}

	return node, parentNode, err
}

// Print - prints the node in json
func (n Node) Print() {
	printTree, _ := json.MarshalIndent(n, "", "|  ")
	fmt.Println(string(printTree))
}
