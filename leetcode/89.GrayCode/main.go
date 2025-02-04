package main

import (
	"fmt"
)

func grayCode(n int) []int {
	result := []int{0}
	for i := 1; i <= n; i++ {
		previousSequenceLength := len(result)
		mask := 1 << (i - 1)
		for j := previousSequenceLength - 1; j >= 0; j-- {
			result = append(result, mask+result[j])
		}
	}
	fmt.Println(2 ^ 2>>1)
	return result
}

func main() {
	n := 3
	fmt.Println(grayCode(n))
}
