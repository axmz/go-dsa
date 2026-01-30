package main

import "fmt"

func hammingDistance(x int, y int) int {
	o := x ^ y
	return hammingWeight(o)
}

func hammingWeight(num int) int {
	sum := 0
	var mask int = 1
	for i := 0; i < 32; i++ {
		if num&mask != 0 {
			sum++
		}
		mask <<= 1
	}

	return sum
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
