package main

import "fmt"

// hard and tricky. better watch the solution.
func maxCoins(nums []int) int {
	memo := make(map[[2]int]int)
	nums = append([]int{1}, nums...)
	nums = append(nums, 1)

	var dp func(i, j int) int
	dp = func(i, j int) int {
		if i > j {
			return 0
		}

		if v, ok := memo[[2]int{i, j}]; ok {
			return v
		}

		maxWin := 0
		for k := i; k <= j; k++ {
			// this part is difficult to understand:
			// we are trying to burst the balloon at position k as a last balloon.
			// so the coins we get is nums[i-1] * nums[k] * nums[j+1],
			// and the remaining coins we get is dp(i, k-1) + dp(k+1, j)
			// it is as if we intervals that are always 3 balloons wide and nearest neighbours are i-1 and j+1,
			// even if i and j are far apart as numbers from k
			coins := nums[i-1] * nums[k] * nums[j+1]
			remaining := dp(i, k-1) + dp(k+1, j)
			win := coins + remaining
			if win > maxWin {
				maxWin = win
			}
		}

		memo[[2]int{i, j}] = maxWin
		return maxWin
	}

	return dp(1, len(nums)-2)
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
