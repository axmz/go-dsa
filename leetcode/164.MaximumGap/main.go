package main

import "fmt"

func maximumGap(nums []int) int {

	n := len(nums)
	if n < 2 {
		return 0
	}

	maxNum := nums[0]
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}

	for exp := 1; maxNum/exp > 0; exp *= 10 {
		count := [10]int{}
		for _, num := range nums {
			digit := (num / exp) % 10
			count[digit]++
		}

		for i := 1; i < 10; i++ {
			count[i] += count[i-1]
		}

		newArr := make([]int, n)
		for i := n - 1; i >= 0; i-- {
			num := nums[i]
			digit := (num / exp) % 10
			count[digit]--
			newArr[count[digit]] = num
		}
		nums = newArr
	}

	maxGap := 0
	for i := 1; i < n; i++ {
		gap := nums[i] - nums[i-1]
		if gap > maxGap {
			maxGap = gap
		}
	}

	return maxGap

}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
