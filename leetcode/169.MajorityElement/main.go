package main

import "fmt"

// Bit manipulation approach
func majorityElement2(nums []int) int {
	majority := 0
	for i := 0; i < 32; i++ {
		mask := int32(1 << i)
		bitcount := 0
		for _, n := range nums {
			if int32(n)&mask != 0 {
				bitcount++
			}
		}
		if bitcount > len(nums)/2 {
			majority |= int(mask)
		}
	}
	return majority
}

// Boyer-Moore Voting Algorithm
func majorityElement(nums []int) int {
	count := 0
	var candidate int

	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count += 1
		} else {
			count -= 1
		}
	}

	return candidate
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
