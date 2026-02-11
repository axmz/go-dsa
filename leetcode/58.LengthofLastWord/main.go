package main

import "fmt"

func lengthOfLastWord(s string) int {
	n := len(s)
	var space byte = ' '
	l, r := n-1, n-1
	for l >= 0 {
		if s[l] == space && s[r] == space {
			r--
			l--
		} else if s[r] != space {
			if s[l] != space {
				l--
			} else {
				return r - l
			}
		}
	}

	return r - l
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
