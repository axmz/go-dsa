package main

import (
	"fmt"
)

func countBits2(n int) []int {
	res := make([]int, n+1)
	for i := 0; i <= n; i++ {
		res[i] = res[i/2] + i%2
	}
	fmt.Printf("%[1]d - %04[1]b \n", 0)
	fmt.Printf("%[1]d - %04[1]b \n", 1)
	fmt.Printf("%[1]d - %04[1]b \n", 2)
	fmt.Printf("%[1]d - %04[1]b \n", 3)
	fmt.Printf("%[1]d - %04[1]b \n", 4)
	fmt.Printf("%[1]d - %04[1]b \n", 5)
	fmt.Printf("%[1]d - %04[1]b \n", 6)
	fmt.Printf("%[1]d - %04[1]b \n", 7)
	fmt.Printf("%[1]d - %04[1]b \n", 8)
	return res
}

func countBits(n int) []int {
	res := make([]int, n+1)

	for b := 1; b <= n; b <<= 1 {
		for x := 0; x < b && x+b <= n; x++ {
			res[b+x] = res[x] + 1
		}
	}
	fmt.Printf("%[1]d - %04[1]b \n", 0)
	fmt.Printf("%[1]d - %04[1]b \n", 1)
	fmt.Printf("%[1]d - %04[1]b \n", 2)
	fmt.Printf("%[1]d - %04[1]b \n", 3)
	fmt.Printf("%[1]d - %04[1]b \n", 4)
	fmt.Printf("%[1]d - %04[1]b \n", 5)
	fmt.Printf("%[1]d - %04[1]b \n", 6)
	fmt.Printf("%[1]d - %04[1]b \n", 7)
	fmt.Printf("%[1]d - %04[1]b \n", 8)
	return res
}
func main() {
	res := countBits(8)
	fmt.Println(res)
}
