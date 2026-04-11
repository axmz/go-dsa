package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	rows := len(obstacleGrid)
	cols := len(obstacleGrid[0])

	dp := make([][]int, rows)
	for i := range dp {
		dp[i] = make([]int, cols)
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if obstacleGrid[r][c] == 1 {
				dp[r][c] = 0
			} else if r == 0 && c == 0 {
				dp[r][c] = 1
			} else if r == 0 {
				dp[r][c] = dp[r][c-1]
			} else if c == 0 {
				dp[r][c] = dp[r-1][c]
			} else {
				dp[r][c] = dp[r-1][c] + dp[r][c-1]
			}
		}
	}

	return dp[rows-1][cols-1]
}

func uniquePathsWithObstacles2(obstacleGrid [][]int) int {
	rows := len(obstacleGrid)
	cols := len(obstacleGrid[0])

	dp := make([][]int, rows)
	for i := range dp {
		dp[i] = make([]int, cols)
	}

	for c := 0; c < cols; c++ {
		if obstacleGrid[0][c] == 1 {
			break
		}
		dp[0][c] = 1
	}

	for r := 0; r < rows; r++ {
		if obstacleGrid[r][0] == 1 {
			break
		}
		dp[r][0] = 1
	}

	for r := 1; r < rows; r++ {
		for c := 1; c < cols; c++ {
			if obstacleGrid[r][c] == 1 {
				continue
			}
			dp[r][c] = dp[r-1][c] + dp[r][c-1]
		}
	}

	return dp[rows-1][cols-1]
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
