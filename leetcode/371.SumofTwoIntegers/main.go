package main

import (
	"fmt"
)

func getSum(a int, b int) int {
	for b != 0 {
		a, b = a^b, a&b<<1
	}

	return a
}

func main() {
	res := getSum(22, 22)
	fmt.Println(res)
}
