package main

import "fmt"

func isHappy(n int) bool {
	seen := make(map[int]bool)
	for {
		if n == 1 {
			return true
		}
		if seen[n] {
			return false
		}
		seen[n] = true
		sum := 0
		for n > 0 {
			digit := n % 10
			sum += digit * digit
			n /= 10
		}
		n = sum
	}
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
