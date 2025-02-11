package main

import "fmt"

// there is a better algo
func plusOne(digits []int) []int {
	last := len(digits) - 1
	if digits[last] != 9 {
		digits[last]++
		return digits
	}

	carry := 0
	for i := last; i >= 0 && (digits[i] == 9 || carry == 1); i-- {
		if digits[i] == 9 && i == 0 {
			digits[i] = 0
			digits = append([]int{1}, digits...)
		} else if digits[i] == 9 {
			digits[i] = 0
			carry = 1
		} else {
			digits[i]++
			carry = 0
			break
		}
	}

	return digits
}

func main() {
	// nums := []int{4, 8, 8, 0}
	nums := []int{9, 8, 9}
	fmt.Println(plusOne(nums))
}
