package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	rem := x
	res := 0
	for rem != 0 {
		d := rem % 10
		rem /= 10
		res = res*10 + d
		if res > math.MaxInt32 || res < math.MinInt32 {
			return 0
		}
	}

	if x < 0 {
		return -res
	}
	return res
}

func main() {
	x := -123
	fmt.Println(reverse(x))
}
