package main

import "fmt"

// func combinationSum3(k int, n int) [][]int {
// 	res := [][]int{}
// 	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

// 	var backtrack func(nums []int, combination []int)
// 	isValid := func(combination []int) bool {
// 		sum := 0
// 		for _, v := range combination {
// 			sum += v
// 		}

// 		max := 0
// 		for _, v := range combination {
// 			if v > max {
// 				max = v
// 			}
// 		}

// 		rem := k - len(combination)

// 		for i := max + 1; i <= rem+max; i++ {
// 			sum += i
// 		}

// 		return sum <= n
// 	}

// 	isCondition := func(combination []int) bool {
// 		if len(combination) != k {
// 			return false
// 		}

// 		sum := 0
// 		for _, v := range combination {
// 			sum += v
// 		}

// 		return sum == n
// 	}

// 	backtrack = func(nums []int, combination []int) {
// 		if isCondition(combination) {
// 			res = append(res, combination)
// 			return
// 		}

// 		if len(nums) == 0 {
// 			return
// 		}

// 		if !isValid(combination) {
// 			return
// 		}

// 		pop := nums[0]
// 		nums = nums[1:]
// 		c := make([]int, len(combination))
// 		copy(c, combination)

// 		if len(combination) < k {
// 			c = append(combination, pop)
// 		} else {
// 			c[len(c)-1] = pop
// 		}

// 		backtrack(nums, c)
// 	}

// 	backtrack(nums, []int{})

//		return res
//	}

var count = 0

// WITH OPTIMIZATION isValid
func combinationSum3(k int, n int) [][]int {
	res := [][]int{}
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	var backtrack func(nums []int, combination []int)

	isCondition := func(combination []int) bool {
		if len(combination) != k {
			return false
		}

		sum := 0
		for _, v := range combination {
			sum += v
		}

		return sum == n
	}

	isValid := func(combination []int) bool {
		sum := 0
		for _, v := range combination {
			sum += v
		}

		max := 0
		for _, v := range combination {
			if v > max {
				max = v
			}
		}

		rem := k - len(combination)

		for i := max + 1; i <= rem+max; i++ {
			sum += i
		}

		return sum <= n
	}

	backtrack = func(nums []int, combination []int) {
		count++
		fmt.Println(count)
		if isCondition(combination) {
			res = append(res, combination)
			return
		}

		if len(nums) == 0 || len(combination) == k {
			return
		}

		for len(nums) > 0 {
			pop := nums[0]
			nums = nums[1:]
			c := make([]int, len(combination))
			copy(c, combination)
			c = append(c, pop)

			if !isValid(c) {
				return
			}

			backtrack(nums, c)
		}
	}

	backtrack(nums, []int{})

	return res
}

func main() {
	// k := 3
	// n := 7
	k := 4
	n := 24
	// k := 3
	// n := 9
	fmt.Println(combinationSum3(k, n))
}
