package main

import "fmt"

func majorityElement(nums []int) []int {
	count1, count2, candidate1, candidate2 := 0, 0, 0, 0
	for _, n := range nums {
		if candidate1 == n {
			count1++
		} else if candidate2 == n {
			count2++
		} else if count1 == 0 {
			candidate1, count1 = n, 1
		} else if count2 == 0 {
			candidate2, count2 = n, 1
		} else {
			count1--
			count2--
		}
	}

	floor := len(nums) / 3
	count1, count2 = 0, 0
	for _, n := range nums {
		switch n {
		case candidate1:
			count1++
		case candidate2:
			count2++
		}
	}

	ans := []int{}
	if count1 > floor {
		ans = append(ans, candidate1)
	}
	if count2 > floor {
		ans = append(ans, candidate2)
	}
	return ans
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
