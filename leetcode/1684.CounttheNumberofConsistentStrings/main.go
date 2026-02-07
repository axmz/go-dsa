package main

import "fmt"

func getBitset(s []byte) uint32 {
	var bitset uint32

	for _, b := range s {
		bitset |= 1 << (b - 'a')
	}

	return bitset
}

func countConsistentStrings(allowed string, words []string) int {
	allowed_bitset := getBitset([]byte(allowed))
	count := 0

	for _, w := range words {
		word_bitset := getBitset([]byte(w))
		if allowed_bitset&word_bitset == word_bitset {
			count++
		}
	}

	return count
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
