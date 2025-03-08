package main

import "fmt"

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	var recurse func(l, r int) *TreeNode
	recurse = func(l, r int) *TreeNode {
		if l > r {
			return nil
		}
		m := (l + r) / 2
		n := &TreeNode{
			Val:   nums[m],
			Left:  recurse(l, m-1),
			Right: recurse(m+1, r),
		}

		return n
	}

	return recurse(0, len(nums)-1)
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
