package main

import (
	"fmt"
)

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

func minDifficultyTopDown(jobDifficulty []int, days int) int {
	n := len(jobDifficulty)
	if n < days {
		return -1
	}

	memo := make([][]int, days+1)
	for i := range memo {
		s := make([]int, n)
		for i := range s {
			s[i] = -1
		}
		memo[i] = s
	}

	remainingDifficulty := make([]int, n)
	localMax := 0
	for i := n - 1; i >= 0; i-- {
		if jobDifficulty[i] > localMax {
			localMax = jobDifficulty[i]
		}

		remainingDifficulty[i] = localMax
	}

	var dp func(d, i int) int
	dp = func(d, i int) int {

		if v := memo[d][i]; v != -1 {
			return v
		}

		if d == days {
			memo[d][i] = remainingDifficulty[i]
			return remainingDifficulty[i]
		}

		j := len(jobDifficulty) - (days - d)

		if i > len(jobDifficulty) {
			return 0
		}

		newmax := -1
		minimal := 1000000
		for k := i; k < j; k++ {
			newmax = max(jobDifficulty[k], newmax)
			res := newmax + dp(d+1, k+1)
			minimal = min(res, minimal)
			memo[d][i] = minimal
		}

		return minimal
	}

	return dp(1, 0)
}

func minDifficultyBottomToTop(jobDifficulty []int, days int) int {
	n := len(jobDifficulty)
	if n < days {
		return -1
	}

	cur, prev := make([]int, n), make([]int, n)

	curMax := 0
	for i := len(prev) - 1; i >= 0; i-- {
		curMax = max(curMax, jobDifficulty[i])
		prev[i] = curMax
	}

	for d := days - 1; d > 0; d-- {
		ilim := n - 1 - (days - d)
		for i := d - 1; i <= ilim; i++ {
			curMax, curMin := 0, 1000000
			for j := i; j <= ilim; j++ {
				curMax = max(curMax, jobDifficulty[j])
				curMin = min(curMin, curMax+prev[j+1])
			}
			cur[i] = curMin
		}

		cur, prev = prev, cur
	}

	return prev[0]
}

// monotonic stack
func minDifficulty(jobDifficulty []int, d int) int {
	n := len(jobDifficulty)
	if n < d {
		return -1
	}

	minDiffPrevDay := make([]int, n)
	minDiffCurrDay := make([]int, n)

	// Initialize minDiffPrevDay with a large value
	for i := range minDiffPrevDay {
		minDiffPrevDay[i] = 1000
	}

	stack := []int{}

	for day := 0; day < d; day++ {
		// Clear the stack for the new day
		stack = stack[:0]

		for i := day; i < n; i++ {
			// Initialize minDiffCurrDay[i]
			if i > 0 {
				minDiffCurrDay[i] = minDiffPrevDay[i-1] + jobDifficulty[i]
			} else {
				minDiffCurrDay[i] = jobDifficulty[i]
			}

			// Maintain a monotonically decreasing stack
			for len(stack) > 0 && jobDifficulty[stack[len(stack)-1]] <= jobDifficulty[i] {
				j := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				diffIncr := jobDifficulty[i] - jobDifficulty[j]
				minDiffCurrDay[i] = min(minDiffCurrDay[i], minDiffCurrDay[j]+diffIncr)
			}

			// If stack is not empty, consider the top of the stack
			if len(stack) > 0 {
				minDiffCurrDay[i] = min(minDiffCurrDay[i], minDiffCurrDay[stack[len(stack)-1]])
			}

			// Push the current index onto the stack
			stack = append(stack, i)
		}

		// Swap current day and previous day arrays
		minDiffPrevDay, minDiffCurrDay = minDiffCurrDay, minDiffPrevDay
	}

	return minDiffPrevDay[n-1]
}

func main() {
	// jobDifficulty := []int{6, 5, 4, 3, 2, 1}
	// d := 2
	jobDifficulty := []int{6, 5, 10, 3, 2, 1}
	d := 3
	// jobDifficulty := []int{7, 1, 7, 1, 7, 1}
	// d := 3
	// jobDifficulty := []int{11, 111, 22, 222, 33, 333, 44, 444}
	// d := 6
	fmt.Println(minDifficulty(jobDifficulty, d))
}
