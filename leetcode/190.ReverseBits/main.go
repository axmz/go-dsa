package main

import (
	"fmt"
)

func reverseBits(num uint32) uint32 {
	var res uint32

	for i := 31; i >= 0; i-- {
		res = res | (num&1)<<i
		num >>= 1
	}

	return res
}

func main() {
	res := reverseBits(22)
	fmt.Println(res)
}
