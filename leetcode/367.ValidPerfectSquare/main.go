package main

import "fmt"

func sqrt(num int) int {
	if num == 0 {
		return 0
	}
	x := num
	for x > num/x {
		x = (x + num/x) / 2
	}
	return x
}

func isPerfectSquare(num int) bool {
	s := sqrt(num)
	return s*s == num
}

func main() {
	x := 5
	fmt.Println(isPerfectSquare(x))
}
