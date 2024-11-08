package main

import "fmt"

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n < 0 {
		n = -n
		x = 1 / x
	}

	if n%2 == 0 {
		return myPow(x*x, n/2)
	}

	return x * myPow(x*x, (n-1)/2)
}

func main() {
	x := 3.0
	n := 0
	fmt.Println(myPow(x, n))

	x2 := 2.0
	n2 := -3
	fmt.Println(myPow(x2, n2))
}
