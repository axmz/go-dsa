package main

import (
	"fmt"
	"math"
	"strconv"
)

func isPowerOfThree2(n int) bool {
	if n <= 0 {
		return false
	}
	m := float64(n)
	log := math.Log10(m) / math.Log10(3)
	epsilon := 1e-10
	return math.Abs(log-math.Round(log)) < epsilon
}

// 1162261467 is the max number that fit int32 that is 3^19
// anything lower than that, divisible by 3 is the power of 3
func isPowerOfThree(n int) bool {
	return n > 0 && 1162261467%n == 0
}

func isPowerOfThree3(n int) bool {
	if n == 1 {
		return true
	}
	if n < 3 {
		return false
	}
	for n%3 == 0 {
		n /= 3
	}
	return n == 1
}

func main() {
	i, _ := strconv.ParseInt("1010", 3, 64) // numbers can be expressed in base other than 10 or 2 or 8
	fmt.Println(i)

	x := 27
	fmt.Println(isPowerOfThree(x))
}
