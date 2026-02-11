package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// TLE
func maxSubArrayTLE(nums []int) int {
	l := len(nums)
	sum := -100000
	for i := 0; i < l; i++ {
		cur := 0
		for j := i; j < l; j++ {
			cur += nums[j]
			sum = max(sum, cur)
		}
	}
	return sum
}

// Kadane
func maxSubArray2(nums []int) int {
	l := len(nums)
	sum := nums[0]
	cur := nums[0]
	for i := 1; i < l; i++ {
		cur = max(nums[i], cur+nums[i])
		sum = max(sum, cur)
	}
	return sum
}

// DAQ
func maxSubArray(nums []int) int {
	var helper func(left, right int) int
	helper = func(left, right int) int {
		if left > right {
			return -10000000
		}

		prev := 0
		maxL := 0
		maxR := 0
		mid := left + (right-left)/2
		for i := mid - 1; i >= left; i-- {
			prev += nums[i]
			maxL = max(maxL, prev)
		}

		prev = 0
		for i := mid + 1; i <= right; i++ {
			prev += nums[i]
			maxR = max(maxR, prev)
		}

		best := maxL + nums[mid] + maxR
		leftSubarraySum := helper(left, mid-1)
		rightSubarraySum := helper(mid+1, right)

		if leftSubarraySum > best {
			best = leftSubarraySum
		}
		if rightSubarraySum > best {
			best = rightSubarraySum
		}
		return best
	}
	return helper(0, len(nums)-1)
}

// DAQ, cleaner code
func maxSubArray4(nums []int) int {
	var helper func(left, right int) int
	helper = func(left, right int) int {
		if left == right {
			return nums[left]
		}
		mid := (left + right) / 2
		leftSum := helper(left, mid)
		rightSum := helper(mid+1, right)

		cur := 0
		leftMax := -100000
		for i := mid; i >= left; i-- {
			cur += nums[i]
			leftMax = max(leftMax, cur)
		}

		cur = 0
		rightMax := -100000
		for i := mid + 1; i <= right; i++ {
			cur += nums[i]
			rightMax = max(rightMax, cur)
		}

		return max(max(leftSum, rightSum), leftMax+rightMax)
	}
	return helper(0, len(nums)-1)
}

func main() {
	// nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	// nums := []int{5, 4, -1, 7, 8}
	nums := []int{1}
	fmt.Println(maxSubArray(nums))
}
