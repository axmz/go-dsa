package main

import "fmt"

func minimumFlips(n int) int {
	r := 0
	originalN := n
	for n > 0 {
		n >>= 1
		r++
	}

	ans := 0
	for l := 0; l < r; l++ {
		if originalN>>l&1 != originalN>>(r-1)&1 {
			ans++
		}
		r--
	}

	return ans * 2
}

func main() {
	x := 10
	fmt.Println(minimumFlips(x))
}
