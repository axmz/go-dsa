package main

import (
	"fmt"
	"math"
)

// This function can be cleaned and optimized further
func divide(dividend int, divisor int) int {
	if math.MinInt32 == dividend && -1 == divisor {
		return math.MaxInt32
	}
	negatives := 2
	if dividend > 0 {
		negatives--
		dividend = -dividend
	}
	if divisor > 0 {
		negatives--
		divisor = -divisor
	}
	quotient := 0
	powerOfTwo := 1
	cumulativeDivisor := divisor

	for divisor >= dividend {
		for (cumulativeDivisor << 1) >= dividend {
			powerOfTwo <<= 1
			cumulativeDivisor <<= 1
		}
		quotient += powerOfTwo
		dividend = dividend - cumulativeDivisor
		powerOfTwo = 1
		cumulativeDivisor = divisor
	}

	if negatives != 1 {
		return quotient
	}

	return -quotient
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
