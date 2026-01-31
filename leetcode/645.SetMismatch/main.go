package main

import "fmt"

// for nums[1 2 4 4 5 6]
// original[1 2 3 4 5 6]
func findErrorNums(nums []int) []int {
	// isolate xor of missing and duplicate number
	xor_pair := 0
	for _, num := range nums {
		xor_pair ^= num
	}
	for i := 1; i <= len(nums); i++ {
		xor_pair ^= i
	}
	// this produced xor_pair = missing ^ duplicate = 3 ^ 4

	// find rightmost set bit that will be used to partition numbers into two groups
	rightmost_set_bit := xor_pair & -xor_pair // 011 ^ 100 = 111

	xor_group_0 := 0 // group with rightmost set bit 0
	xor_group_1 := 0 // group with rightmost set bit 1

	// this separates into
	// 1	001
	// 2		010
	// 4		100
	// 4		100
	// 5	101
	// 6 		110
	for _, num := range nums {
		if num&rightmost_set_bit == 0 {
			xor_group_0 ^= num
		} else {
			xor_group_1 ^= num
		}
	}

	// this separates into
	// 1	001
	// 2		010
	// 3	011
	// 4		100
	// 5	101
	// 6 		110
	// and after xoring the group on the left and on the right it isolates missing and duplicate
	for i := 1; i <= len(nums); i++ {
		if i&rightmost_set_bit == 0 {
			xor_group_0 ^= i
		} else {
			xor_group_1 ^= i
		}
	}

	// sort into required order
	for _, num := range nums {
		if num == xor_group_0 {
			return []int{xor_group_0, xor_group_1}
		}
	}
	return []int{xor_group_1, xor_group_0}
}

func main() {
	nums := []int{1, 2, 4, 4, 5, 6}
	fmt.Println(findErrorNums(nums))
}
