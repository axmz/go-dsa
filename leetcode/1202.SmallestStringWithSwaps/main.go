package main

import (
	"fmt"
	. "godsa/utils/disjointset"
	"sort"
)

func smallestStringWithSwaps(s string, pairs [][]int) string {
	u := NewUnion(len(s))
	for _, e := range pairs {
		u.Union(e[0], e[1])
	}

	m := make(map[int][]int)
	for i := 0; i < len(s); i++ {
		root := u.Find(i)
		m[root] = append(m[root], i)
	}

	// Result array
	res := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		res[i] = s[i] // Initialize with original string
	}

	// Process each component
	for _, indices := range m {
		// Extract characters
		chars := make([]byte, len(indices))
		for i, idx := range indices {
			chars[i] = s[idx]
		}
		// Sort characters
		sort.Slice(chars, func(i, j int) bool {
			return chars[i] < chars[j]
		})
		// Sort indices to align with result positions
		sort.Ints(indices)
		// Place sorted characters in original positions
		for i, idx := range indices {
			res[idx] = chars[i]
		}
	}

	return string(res)
}

func main() {
	// s := "dcxxxba"
	// nums := [][]int{{0, 3}, {1, 2}, {0, 2}}
	s := "dcab"
	nums := [][]int{{0, 3}, {1, 2}}
	fmt.Println(smallestStringWithSwaps(s, nums))
}
