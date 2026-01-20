package main

import (
	"fmt"
	"math/rand"
)

type Node struct {
	Val      int
	Priority int // Heap property: parent priority > child priority
	Left     *Node
	Right    *Node
}

type Treap struct {
	root *Node
}

func NewTreap() *Treap {
	return &Treap{root: nil}
}

// Right rotation: used when left child has higher priority
//
//	   y                x
//	  / \              / \
//	 x   C    ===>    A   y
//	/ \                  / \
//
// A   B                B   C
func (t *Treap) rotateRight(y *Node) *Node {
	x := y.Left
	y.Left = x.Right
	x.Right = y
	return x
}

// Left rotation: used when right child has higher priority
//
//	 x                  y
//	/ \                / \
//
// A   y      ===>    x   C
//
//	 / \            / \
//	B   C          A   B
func (t *Treap) rotateLeft(x *Node) *Node {
	y := x.Right
	x.Right = y.Left
	y.Left = x
	return y
}

func (t *Treap) Insert(val int) {
	t.root = t.insert(t.root, val)
}

// Recursive insert: insert as BST, then rotate to maintain heap property
func (t *Treap) insert(node *Node, val int) *Node {
	if node == nil {
		return &Node{
			Val:      val,
			Priority: rand.Intn(1 << 30), // Random priority
		}
	}

	if val < node.Val {
		node.Left = t.insert(node.Left, val)
		// If left child has higher priority, rotate right
		if node.Left.Priority > node.Priority {
			node = t.rotateRight(node)
		}
	} else {
		node.Right = t.insert(node.Right, val)
		// If right child has higher priority, rotate left
		if node.Right.Priority > node.Priority {
			node = t.rotateLeft(node)
		}
	}

	return node
}

func (t *Treap) Delete(val int) {
	t.root = t.delete(t.root, val)
}

// Delete by rotating node down to leaf, then removing it
func (t *Treap) delete(node *Node, val int) *Node {
	if node == nil {
		return nil
	}

	if val < node.Val {
		node.Left = t.delete(node.Left, val)
	} else if val > node.Val {
		node.Right = t.delete(node.Right, val)
	} else {
		// Found the node to delete
		// Rotate it down until it becomes a leaf
		if node.Left == nil && node.Right == nil {
			return nil // Leaf node, remove it
		} else if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		} else {
			// Two children: rotate the child with higher priority up
			if node.Left.Priority > node.Right.Priority {
				node = t.rotateRight(node)
				node.Right = t.delete(node.Right, val)
			} else {
				node = t.rotateLeft(node)
				node.Left = t.delete(node.Left, val)
			}
		}
	}

	return node
}

// LowerBound finds the smallest element >= val
func (t *Treap) LowerBound(val int) *Node {
	var result *Node
	current := t.root

	for current != nil {
		if current.Val >= val {
			result = current
			current = current.Left
		} else {
			current = current.Right
		}
	}
	return result
}

func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
	treap := NewTreap()
	for i := 0; i < len(nums); i++ {
		if i > indexDiff {
			treap.Delete(nums[i-indexDiff-1])
		}

		// Find lower bound of (nums[i] - valueDiff)
		// This gives us the smallest element >= (nums[i] - valueDiff)
		// Which is basically the first/closest candidate that could satisfy the condition
		lb := treap.LowerBound(nums[i] - valueDiff)

		// If lb exists and is within range [nums[i] - valueDiff, nums[i] + valueDiff]
		if lb != nil && lb.Val <= nums[i]+valueDiff {
			return true
		}

		treap.Insert(nums[i])
	}
	return false
}

func main() {
	indexDiff := 2
	valueDiff := 3
	nums := []int{1, 5, 9, 1, 5, 9}

	fmt.Println("Lower bound approach:", containsNearbyAlmostDuplicate(nums, indexDiff, valueDiff))
}
