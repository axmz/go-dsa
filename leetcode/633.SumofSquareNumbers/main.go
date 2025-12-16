package main

import (
	"fmt"
	"math"
)

// func judgeSquareSum(c int) bool {
// 	for b := 1; b*b <= c; b++ {
// 		diff := c - b*b
// 		a := math.Sqrt(float64(diff))
// 		if a == float64(int(a)) {
// 			return true
// 		}
// 	}

// 	return false
// }

func judgeSquareSum(c int) bool {
	a := 0
	b := int(math.Sqrt(float64(c)))

	for a <= b {
		r := a*a + b*b
		if r == c {
			return true
		} else if r < c {
			a++
		} else {
			b--
		}
	}

	return false
}

func main() {
	x := 26
	fmt.Printf("%t\n", judgeSquareSum(x))
}
