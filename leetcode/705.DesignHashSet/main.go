package main

import "fmt"

type MyHashSet struct {
	buckets []*Node
	size    int
}

func Constructor() MyHashSet {
	return MyHashSet{buckets: make([]*Node, 769), size: 769}
}

func (this *MyHashSet) _hash(key int) int {
	return key % this.size
}

func (this *MyHashSet) Add(key int) {
	h := this._hash(key)
	if this.buckets[h] == nil {
		this.buckets[h] = NewNode(key)
		return
	}
	this.buckets[h].Insert(key)
}

func (this *MyHashSet) Remove(key int) {
	h := this._hash(key)
	if this.buckets[h] != nil {
		this.buckets[h] = this.buckets[h].Remove(key)
	}
}

func (this *MyHashSet) Contains(key int) bool {
	h := this._hash(key)
	if this.buckets[h] == nil {
		return false
	}
	if _, _, err := this.buckets[h].Lookup(key, nil); err != nil {
		return false
	}
	return true
}

// Node structure of the node
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func NewNode(value int) *Node {
	return &Node{Value: value}
}

// Insert inserts a value into BST
func (n *Node) Insert(num int) {
	if num == n.Value {
		return
	}
	if num > n.Value {
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

// Lookup searches for a value in BST and returns the node and its parent
func (n *Node) Lookup(num int, parent *Node) (*Node, *Node, error) {
	if n == nil {
		return nil, parent, fmt.Errorf("node not found")
	}
	if n.Value == num {
		return n, parent, nil
	}
	if num > n.Value {
		return n.Right.Lookup(num, n)
	}
	return n.Left.Lookup(num, n)
}

// Remove removes the node with the given value from the tree and returns the new root
func (n *Node) Remove(num int) *Node {
	if n == nil {
		return nil
	}

	if num < n.Value {
		n.Left = n.Left.Remove(num)
		return n
	} else if num > n.Value {
		n.Right = n.Right.Remove(num)
		return n
	}

	// Found the node to remove
	// Case 1: No children
	if n.Left == nil && n.Right == nil {
		return nil
	}
	// Case 2: Only right child
	if n.Left == nil {
		return n.Right
	}
	// Case 3: Only left child
	if n.Right == nil {
		return n.Left
	}

	// Case 4: Two children - find inorder successor (smallest in right subtree)
	minRight := n.Right
	for minRight.Left != nil {
		minRight = minRight.Left
	}
	// Replace current node's value with successor's value
	n.Value = minRight.Value
	// Remove the successor node
	n.Right = n.Right.Remove(minRight.Value)
	return n
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
