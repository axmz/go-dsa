package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
const (
	offset    uint64 = 0xcbf29ce484222325 // FNV offset basis
	prime     uint64 = 0x100000001b3      // FNV prime
	tagNil    uint64 = 0xdeadbeefcafebabe // not good for cryptography
	saltLeft  uint64 = 0x9E3779B97F4A7C15 // 2^64 / golden ratio
	saltRight uint64 = 0xA0761D6478BD642F // Comes from SplitMix64 / Wyhash family
)

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	m := map[uint64]*TreeNode{}
	seen := map[uint64]bool{}
	res := []*TreeNode{}

	var fnv func(node *TreeNode) uint64
	fnv = func(node *TreeNode) uint64 {
		if node == nil {
			return tagNil
		}
		h := offset
		h ^= uint64(node.Val)
		h *= prime                     // * has better mixing properties than +
		h ^= fnv(node.Left) ^ saltLeft // note: + or ^ operations have slightly different properties
		h *= prime
		h ^= fnv(node.Right) ^ saltRight
		h *= prime

		if n, exists := m[h]; exists {
			if !seen[h] {
				res = append(res, n)
				seen[h] = true
			}
		} else {
			m[h] = node
		}

		return h
	}

	fnv(root)

	return res
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
