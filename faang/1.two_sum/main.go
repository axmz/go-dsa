package main

import "fmt"

func main() {
	arr := []int{1, 3, 7, 9, 2}
	sum := 11
	var result [2]int
	m := make(map[int]int)
	for i, num := range arr {
		sumPair, sumPairFound := m[num]
		if sumPairFound {
			result[0] = i
			result[1] = sumPair
		} else {
			dif := sum - num
			m[dif] = i
		}
	}
	fmt.Println(result)
	result[0] = 0
	result[1] = 0
}
