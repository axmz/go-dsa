package main

import "fmt"

func solution(knows func(a int, b int) bool) func(n int) int {

	return func(n int) int {
		candidate := 0
		for i := 1; i < n; i++ {
			if knows(candidate, i) {
				candidate = i
			}
		}

		for i := 0; i < n; i++ {
			if i == candidate {
				continue
			}
			// -1 if candidate knows anyone or if anyone doesn't know candidate
			if knows(candidate, i) || !knows(i, candidate) {
				return -1
			}
		}

		return candidate
	}
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
