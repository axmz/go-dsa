package main

import (
	"fmt"
	. "godsa/utils/disjointset"
)

func factorization(n int) []int {
	factors := make([]int, 0)
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			factors = append(factors, i)
			n /= i
		}
	}

	if n > 1 {
		factors = append(factors, n)
	}

	return factors
}

func canTraverseAllPairs(nums []int) bool {
	n := len(nums)
	u := NewUnion(n)
	factorMap := make(map[int][]int)

	for i, num := range nums {
		factors := factorization(num)
		for _, f := range factors {
			factorMap[f] = append(factorMap[f], i)
		}
	}

	for _, indices := range factorMap {
		for i := 1; i < len(indices); i++ {
			u.Union(indices[0], indices[i])
		}
	}

	root := u.Find(0)
	for i := 1; i < n; i++ {
		if u.Find(i) != root {
			return false
		}
	}

	return true
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
