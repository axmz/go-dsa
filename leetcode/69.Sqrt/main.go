package main

import "fmt"

func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	var sq int
	mid, l, r := 2, 2, x/2
	for l <= r {
		mid = l + (r-l)>>1 // this instead of (l + r) / 2 to prevent overflow
		sq = mid * mid
		if sq > x {
			r = mid - 1
		} else if sq < x {
			l = mid + 1
		} else {
			return mid
		}
	}
	return r
}

func main() {
	x := 4
	fmt.Println(mySqrt(x))
}
