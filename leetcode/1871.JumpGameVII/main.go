package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func canReach(s string, minJump int, maxJump int) bool {
	if len(s) == 0 || s[0] != '0' {
		return false
	}

	farthest := 0
	queue := []int{0}

	for len(queue) > 0 {
		i := queue[0]
		queue = queue[1:]
		if i == len(s)-1 {
			return true
		}

		start := max(i+minJump, farthest+1)
		end := min(i+maxJump, len(s)-1)
		for j := start; j <= end; j++ {
			if s[j] == '0' {
				queue = append(queue, j)
			}
		}

		farthest = max(farthest, end)
	}
	return false
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
