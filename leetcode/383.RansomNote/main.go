package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	m := make(map[byte]int)
	for i := 0; i < len(magazine); i++ {
		m[magazine[i]]++
	}

	for i := 0; i < len(ransomNote); i++ {
		if b, ok := m[ransomNote[i]]; !ok || b == 0 {
			return false
		} else {
			m[ransomNote[i]]--
		}
	}
	return true
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
