package main

import (
	"fmt"
)

func toHex(num int) string {
	if num == 0 {
		return "0"
	}

	const (
		hex  = "0123456789abcdef"
		mask = 0xf
	)

	n := uint32(num)
	res := ""

	for n != 0 {
		res = string(hex[n&mask]) + res
		n >>= 4
	}

	return res
}

func main() {
	// num := 1012
	// num := -10
	num := -7
	// num := 0
	fmt.Println(toHex(num))
}
