package main

import "fmt"

func bitwiseComplement(num int) int {
	if num == 0 {
		return 1
	}
	mask := 1
	n := num
	for n > 0 {
		n >>= 1
		mask <<= 1
	}
	return ^num & (mask - 1)
}

func main() {
	x := 5
	res := bitwiseComplement(x)
	fmt.Println(res)
	fmt.Printf("%08b\n", res)
}
