package main

import "fmt"

func modExp(base, exp, mod int64) int64 {
	r := int64(1)
	b := base % mod

	for exp > 0 {
		if exp&1 == 1 {
			r = (r * b) % mod
		}
		b = (b * b) % mod
		exp >>= 1
	}

	return r
}

func countGoodNumbers(n int64) int {
	if n == 1 {
		return 5
	}
	var mod int64 = 1e9 + 7
	var PRIMES int64 = 4 // 2, 3, 5, 7
	var ODD int64 = 5    // 0, 2, 4, 6, 8

	evenCount := (n + 1) / 2 // cool trick. otherwise we would have to have two branches: n%2==0 and n%2==1
	oddCount := n / 2
	return int((modExp(ODD, evenCount, mod) * modExp(PRIMES, oddCount, mod)) % mod)
}

func main() {
	var x int64 = 4
	fmt.Println(countGoodNumbers(x))
}
