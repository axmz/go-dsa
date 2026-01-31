package main

import "fmt"

func hasAlternatingBits(n int) bool {
	lsb := n & 1
	for n > 0 {
		if n&1 != lsb {
			return false
		}
		n >>= 1
		lsb = ^lsb & 1
	}

	return true
}

func main() {
	fmt.Println(hasAlternatingBits(3))
}
