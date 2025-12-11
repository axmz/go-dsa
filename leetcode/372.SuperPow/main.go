package main

import (
	"fmt"
)

func superPow(a int, b []int) int {
	const mod = 1337
	res := 1
	for _, digit := range b {
		// res = (res^10 * a^digit) % mod
		res = (modExp(res, 10, mod) * modExp(a, digit, mod)) % mod
	}
	return res
}

func modExp(x int, n int, p int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x % p
	}
	if n%2 == 0 {
		return modExp((x*x)%p, n/2, p)
	}
	return (x * modExp((x*x)%p, n/2, p)) % p
}

func main() {
	a := 3
	small := []int{3, 2, 3, 3}
	large := []int{4, 3, 3, 8, 5, 2}

	r2 := superPow(a, small)
	fmt.Println(r2)
	r2 = superPow(a, large)
	fmt.Println(r2)
}
