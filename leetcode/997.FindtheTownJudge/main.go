package main

import "fmt"

func findJudge(n int, trust [][]int) int {
	t := make([]int, n+1)
	for _, pair := range trust {
		t[pair[0]]--
		t[pair[1]]++
	}

	for i := 1; i <= n; i++ {
		if t[i] == n-1 {
			return i
		}
	}

	return -1
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
