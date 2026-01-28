package main

import "fmt"

func numJewelsInStones(jewels string, stones string) int {
	j := make(map[byte]bool)
	for i := 0; i < len(jewels); i++ {
		j[jewels[i]] = true
	}
	count := 0
	for i := 0; i < len(stones); i++ {
		if j[stones[i]] {
			count++
		}
	}
	return count
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
