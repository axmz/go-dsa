package main

import "fmt"

func similarPairs(words []string) int {
	res := 0
	hs := make(map[uint32]int)
	for _, w := range words {
		var bitset uint32 = 0
		for i := 0; i < len(w); i++ {
			b := w[i] - 'a'
			bitset |= (1 << b)
		}
		// updating result while in progress, is a good technique to remember
		// otherwise you need another loop of extrawork for counting, sorting, etc.
		// in out case it would be this:
		// for _, x := range hs {
		// 	res += (x - 1) * x / 2
		// }
		count := hs[bitset]
		// note: how combinations were calculated here
		res += count
		hs[bitset] = count + 1
	}
	return res
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
