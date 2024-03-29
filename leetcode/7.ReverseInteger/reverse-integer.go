package main

import (
	"fmt"
	"math"
)

func reverseInteger(x int) int {
	var rev int = 0

	for x != 0 {
		pop := x % 10
		x = x / 10

		if rev > math.MaxInt32/10 || (rev == math.MaxInt32 && pop > 7) {
			return 0
		}

		if rev < math.MinInt32/10 || (rev == math.MinInt32 && pop < -8) {
			return 0
		}

		rev = rev*10 + pop
	}

	return rev
}

func main() {
	fmt.Println("Reverse Integer:", reverseInteger(-123))
}
