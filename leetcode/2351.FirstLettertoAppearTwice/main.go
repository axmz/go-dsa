package main

import "fmt"

func repeatedCharacter(s string) byte {
	var bitset uint32 = 0 // when we deal with alphabets uint32 is enough
	for i := 0; i < len(s); i++ {
		b := s[i] - 'a'
		if (bitset & (1 << b)) != 0 {
			return s[i]
		}
		bitset |= (1 << b)
	}
	return 0
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
