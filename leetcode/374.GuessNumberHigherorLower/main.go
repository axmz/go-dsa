package main

import (
	"fmt"
)

func guess(n int) int {
	pick := 10

	if n > pick {
		return -1
	}

	if n < pick {
		return 1
	}

	return 0
}

func guessNumberBinary(n int) int {
	l, r := 0, n+1
	g := (l + r) / 2

	for res := guess(g); res != 0; res = guess(g) {
		if res == 0 {
			return g
		} else if res == 1 {
			l = g
		} else if res == -1 {
			r = g
		}
		g = (l + r) / 2
	}

	return g
}

func guessNumberTernary(n int) int {
	low := 1
	high := n
	for low <= high {
		mid1 := low + (high-low)/3
		mid2 := high - (high-low)/3
		res1 := guess(mid1)
		res2 := guess(mid2)
		if res1 == 0 {
			return mid1
		}
		if res2 == 0 {
			return mid2
		} else if res1 < 0 {
			high = mid1 - 1
		} else if res2 > 0 {
			low = mid2 + 1
		} else {
			low = mid1 + 1
			high = mid2 - 1
		}
	}
	return -1
}

func main() {
	input := 10

	res := guessNumberTernary(input)
	fmt.Println(res)
}
