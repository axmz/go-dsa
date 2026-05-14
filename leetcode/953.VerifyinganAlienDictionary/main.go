package main

import "fmt"

func isAlienSorted(words []string, order string) bool {
	orderMap := make(map[byte]int)
	for i := 0; i < len(order); i++ {
		c := order[i]
		orderMap[c] = i
	}

	for i := 1; i < len(words); i++ {
		prev, cur := words[i-1], words[i]
		for j := 0; j < len(prev) && j < len(cur); j++ {
			prev_c, cur_c := prev[j], cur[j]
			if prev_c != cur_c {
				if orderMap[prev_c] > orderMap[cur_c] {
					return false
				}
				break
			}
			// if we reach the end of one word, the shorter word should come first
			if j == len(cur)-1 && len(prev) > len(cur) {
				return false
			}
		}
	}

	return true
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
