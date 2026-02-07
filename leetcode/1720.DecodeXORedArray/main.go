package main

import "fmt"

func decode(encoded []int, first int) []int {
	l := len(encoded)
	res := make([]int, l+1)
	res[0] = first
	for i := 0; i < l; i++ {
		res[i+1] = res[i] ^ encoded[i]
	}
	return res
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
