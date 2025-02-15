package main

import "fmt"

func firstUniqChar(s string) int {
	alphabet := [26]int{}
	for i := 0; i < len(s); i++ {
		alphabet[s[i]-'a']++
	}

	for i := 0; i < len(s); i++ {
		if alphabet[s[i]-'a'] == 1 {
			return i
		}
	}

	return -1
}

func main() {
	s := "leetcode"
	fmt.Println(firstUniqChar(s))
}
