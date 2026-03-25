package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	s1hash := [26]byte{}
	s2hash := [26]byte{}
	for i := 0; i < len(s1); i++ {
		s1hash[s1[i]-'a']++
		s2hash[s2[i]-'a']++
	}

	for i := len(s1); i < len(s2); i++ {
		if s1hash == s2hash {
			return true
		}
		s2hash[s2[i]-'a']++
		s2hash[s2[i-len(s1)]-'a']--
	}
	return s1hash == s2hash
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
