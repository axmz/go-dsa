package main

import (
	"fmt"
)

func hammingWeight2(num uint32) int {
	sum := 0
	// fmt.Printf("%032b \n", num)
	// fmt.Printf("%032b \n", num-1)
	// num &= num - 1
	// fmt.Printf("%032b \n", num)
	// fmt.Printf("%032b \n", num-1)
	// num &= num - 1
	// fmt.Printf("%032b \n", num)
	for num != 0 {
		fmt.Println(1)
		sum++
		num &= num - 1
	}

	return sum
}

func hammingWeight(num uint32) int {
	sum := 0
	var mask uint32 = 1
	for i := 0; i < 32; i++ {
		if num&mask != 0 {
			sum++
		}
		mask <<= 1
	}

	return sum
}

func main() {
	res := hammingWeight(22)
	// hammingWeight(9)
	fmt.Println(res)
}
