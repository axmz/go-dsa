package main

import (
	"fmt"
	"math"
)

// Dynamic Programming approach
func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = i // worst case: all 1s
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Breadth-First Search approach
func numSquares2(n int) int {
	q := []int{n}
	visited := make(map[int]bool)
	// visited[n] = true
	steps := 0
	for len(q) > 0 {
		size := len(q)
		steps++
		for i := 0; i < size; i++ {
			pop := q[0]
			q = q[1:]
			for j := 1; j*j <= pop; j++ {
				rem := pop - j*j
				if rem == 0 {
					return steps
				}
				if !visited[rem] {
					visited[rem] = true
					q = append(q, rem)
				}
			}
		}
	}
	return steps
}

func isSquare(n int) bool {
	if n < 0 {
		return false
	}
	root := int(math.Sqrt(float64(n)))
	return root*root == n
}

// Mathematical approach based on Lagrange's Four Square Theorem
func numSquares3(n int) int {
	//Case 1: Result is 1
	if isSquare(n) {
		return 1
	}

	//Case 2: Result is 2 if n can be split into form n = a^2 + b^2
	for i := 1; i*i <= n; i++ {
		if isSquare(n - i*i) {
			return 2
		}
	}

	//Case 4: Result is 4 if n is of the form 4^k(8m + 7)
	temp := n
	for temp%4 == 0 {
		temp /= 4
	}
	if temp%8 == 7 {
		return 4
	}

	//Case 3: Result will be 3 if NOTA
	return 3
}

func main() {
	fmt.Println(numSquares2(12)) // Output: 3 (4+4+4)
	fmt.Println(numSquares2(13)) // Output: 2 (4+9)
}
