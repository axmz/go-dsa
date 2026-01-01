package main

import (
	"fmt"
	"math"
)

func fib2(n int) int {
	cur := 1
	prev := 0
	res := 1

	for i := 0; i < n; i++ {
		res = cur + prev
		prev = cur
		cur = res
	}

	return res
}

func fib(n int) int {
	sqrt5 := math.Sqrt(5)
	goldernRatio := (1 + sqrt5) / 2
	return int(math.Round((math.Pow(goldernRatio, float64(n)) - math.Pow(1-goldernRatio, float64(n))) / sqrt5))
}

func main() {
	x := 10
	fmt.Println(x, fib(x))
}
