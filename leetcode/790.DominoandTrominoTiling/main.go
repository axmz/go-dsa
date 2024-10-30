package main

import "fmt"

func numTilings(n int) int {
	fcache := make(map[int]int)
	pcache := make(map[int]int)

	fcache[1] = 1
	fcache[2] = 2
	pcache[2] = 1

	mod := 1_000_000_007

	var f func(n int) int
	var p func(n int) int

	f = func(n int) int {
		if v, ok := fcache[n]; ok {
			return v
		}

		p(n)
		res := (f(n-1) + f(n-2) + 2*p(n-1)) % mod
		fcache[n] = res
		return res

	}

	p = func(n int) int {
		if v, ok := pcache[n]; ok {
			return v
		}

		res := (p(n-1) + f(n-2)) % mod
		pcache[n] = res
		return res
	}

	return f(n)
}

func main() {
	n := 6
	fmt.Println(numTilings(n))
}
