package main

import "fmt"

// so tricky
func isIsomorphic(s string, t string) bool {
	l := len(s)
	abc := make([]byte, 256)
	seen := make([]bool, 256)
	for i := range abc {
		abc[i] = 0xFF
	}
	for i := 0; i < l; i++ {
		if abc[s[i]] == 0xFF {
			if seen[t[i]] {
				return false
			}
			abc[s[i]] = t[i]
			seen[t[i]] = true
		}
		if abc[s[i]] != t[i] {
			return false
		}
	}
	return true
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
