package main

import (
	"fmt"
	"strings"
)

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func (root *Node) signature() string {
	var sb strings.Builder
	for {
		sb.WriteString(fmt.Sprint(root.Val, ","))
		if root.Next == nil {
			break
		}

		root = root.Next
	}

	return strings.Trim(sb.String(), ",")
}

// func (root *Node) tail() *Node {
// 	cn := root
// 	for cn.Next != nil {
// 		cn = cn.Next
// 	}

// 	return cn
// }

func tail(root *Node) *Node {
	cn := root
	for cn.Next != nil {
		cn = cn.Next
	}

	return cn
}

func flatten(root *Node) *Node {
	if root == nil {
		return nil
	}

	cn := root
	next := cn.Next

	for cn.Next != nil || cn.Child != nil {
		if cn.Child != nil {
			f := flatten(cn.Child)
			f.Prev = cn
			cn.Child = nil
			cn.Next = f

			t := tail(f)
			if next != nil {
				next.Prev, t.Next = t, next
			}
		}

		if next == nil {
			return root
		}

		cn = next
		next = next.Next

	}

	return root
}

func create(list []interface{}) *Node {
	ch := new(Node)
	begin := ch
	cn := ch

	p := new(Node)

	for _, item := range list {
		if n, ok := item.(int); ok { // int
			if *p != (Node{}) { // child
				newNode := &Node{Val: n}
				p.Child, ch, cn = newNode, newNode, newNode
				p = &Node{}
			} else if *ch == (Node{}) { // only once
				ch.Val = n
			} else { // newNode node
				newNode := &Node{Val: n, Prev: cn}
				cn.Next, cn = newNode, newNode
			}
		} else { // nil
			if *p == (Node{}) { // set pointer
				p = ch
			} else { // move pointer
				p = p.Next
			}
		}

	}

	return begin
}

func main() {
	l := []interface{}{1, 2, 3, 4, 5, 6, nil, nil, nil, 7, 8, 9, 10, nil, nil, 11, 12}
	// l := []interface{}{1, 2, nil, 3}
	// l := []interface{}{}
	// l := []interface{}{1, nil, 2, nil, 3, nil}

	ll := create(l)
	f := flatten(ll)
	s := f.signature()
	fmt.Println(s)

}
