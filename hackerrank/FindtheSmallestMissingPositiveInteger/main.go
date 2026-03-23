package main

import "fmt"

func findSmallestMissingPositive(orderNumbers []int32) int32 {
	n := int32(len(orderNumbers))

	// Step 1: Place each number in its correct position
	var i int32 = 0
	for i < n {
		// Calculate the correct position for current number
		correctPos := orderNumbers[i] - 1

		// If the number is positive, within range, and not already in correct position
		if orderNumbers[i] > 0 && orderNumbers[i] <= n && orderNumbers[correctPos] != orderNumbers[i] {
			// Swap the numbers
			orderNumbers[i], orderNumbers[correctPos] = orderNumbers[correctPos], orderNumbers[i]
		} else {
			// Move to next number if we can't place current number
			i++
		}
	}

	// Step 2: Find the first position where number doesn't match index+1
	var j int32 = 0
	for ; j < n; j++ {
		if orderNumbers[j] != j+1 {
			return j + 1
		}
	}

	// If all positions match, the answer is n+1
	return n + 1
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
