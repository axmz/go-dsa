package main

import "fmt"

func kthGrammar(n int, k int) int {
	if n == 1 {
		return 0
	}

	total := 1 << (n - 1) // 2^(n-1)

	if k > total/2 {
		return 1 - kthGrammar(n-1, k-total/2)
	}
	return kthGrammar(n-1, k)
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
