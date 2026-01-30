package main

import "fmt"

// type BitwiseTrie struct {
// 	children [2]*BitwiseTrie
// }

// func Constructor() BitwiseTrie {
// 	return BitwiseTrie{
// 		children: [2]*BitwiseTrie{},
// 	}
// }

// func (this *BitwiseTrie) Insert(num int) {
// 	current := this
// 	for i := 31; i >= 0; i-- {
// 		bit := num >> i & 1
// 		if t := current.children[bit]; t != nil {
// 			current = t
// 		} else {
// 			current.children[bit] = &BitwiseTrie{
// 				children: [2]*BitwiseTrie{},
// 			}
// 			current = current.children[bit]
// 		}
// 	}
// }

// func (this *BitwiseTrie) FindMaxXOR(num int) int {
// 	current := this
// 	maxXOR := 0
// 	for i := 31; i >= 0; i-- {
// 		bit := num >> i & 1
// 		toggleBit := bit ^ 1
// 		if t := current.children[toggleBit]; t != nil {
// 			maxXOR |= (1 << i)
// 			current = t
// 		} else {
// 			current = current.children[bit]
// 		}
// 	}
// 	return maxXOR
// }

// func findMaximumXOR(nums []int) int {
// 	trie := Constructor()
// 	for _, num := range nums {
// 		trie.Insert(num)
// 	}

//		maxXOR := 0
//		for _, num := range nums {
//			xor := trie.FindMaxXOR(num)
//			if xor > maxXOR {
//				maxXOR = xor
//			}
//		}
//		return maxXOR
//	}

// Optimized
type BitwiseTrie struct {
	children [2]*BitwiseTrie
}

func findMaximumXOR(nums []int) int {
	trie := &BitwiseTrie{
		children: [2]*BitwiseTrie{},
	}

	maxXOR := 0
	for _, num := range nums {
		current := trie
		xorNode := trie
		xor := 0
		for i := 31; i >= 0; i-- {
			bit := num >> i & 1
			toggleBit := bit ^ 1
			if t := current.children[bit]; t != nil {
				current = t
			} else {
				current.children[bit] = &BitwiseTrie{
					children: [2]*BitwiseTrie{},
				}
				current = current.children[bit]
			}
			if t := xorNode.children[toggleBit]; t != nil {
				xor |= (1 << i)
				xorNode = t
			} else if xorNode.children[bit] != nil {
				xorNode = xorNode.children[bit]
			}
		}
		if xor > maxXOR {
			maxXOR = xor
		}
	}

	return maxXOR
}

func main() {
	nums := []int{3, 10, 5, 25, 2, 8}
	fmt.Println(findMaximumXOR(nums))
}
