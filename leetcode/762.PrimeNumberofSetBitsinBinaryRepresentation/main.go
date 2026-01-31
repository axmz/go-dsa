package main

import "fmt"

func countPrimeSetBits(left int, right int) int {
	count := 0
	for i := left; i <= right; i++ {
		bitCount := countBits(i)
		if isSmallPrime(bitCount) {
			count++
		}
	}
	return count
}

func countBits(n int) int {
	count := 0
	for n > 0 {
		n &= n - 1
		count++
	}
	return count
}

func isSmallPrime(n int) bool {
	switch n {
	case 2, 3, 5, 7, 11, 13, 17, 19:
		return true
	}
	return false
}

func checkIsPrime(x int) bool {
	if x%2 == 0 || x == 1 {
		return x == 2
	} else {
		for i := 3; i*i <= x; i += 2 {
			if x%i == 0 {
				return false
			}
		}
	}
	return true
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
