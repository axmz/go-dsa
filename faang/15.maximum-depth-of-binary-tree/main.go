package main

import (
	"fmt"
	. "godsa/faang/utils"
)

// [3,9,20,null,null,15,7]
// [5,4,7,3,null,2,null,-1,null,9]
func main() {
	// m := []interface{}{3, 9, 20, nil, nil, 15, 7}
	m := []interface{}{5, 4, 7, 3, nil, 2, nil, -1, nil, 9}
	// m := []interface{}{nil}
	p := CreateTree(m)
	d := MaxDepth(p)
	fmt.Println(p, d)
}
