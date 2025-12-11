package main

import "fmt"

func pow(a int, b int, p int) int {
	r := 1
	base := a % p

	for b > 0 {
		if b&1 == 1 {
			r = (r * base) % p
		}
		base = (base * base) % p
		b >>= 1
	}

	return r
}

func getGoodIndices(variables [][]int, target int) []int {
	result := []int{}
	for i, d := range variables {
		if pow(pow(d[0], d[1], 10), d[2], d[3]) == target {
			result = append(result, i)
		}
	}

	return result
}

func main() {
	x := 2
	// fmt.Println(pow(3, 5, 1000))
	nums := [][]int{{2, 3, 3, 10}, {3, 3, 3, 1}, {6, 1, 1, 4}}
	fmt.Println(getGoodIndices(nums, x))
}
