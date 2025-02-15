package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	alphabet := [26]int{}

	for i := 0; i < len(s); i++ {
		alphabet[s[i]-'a']++
		alphabet[t[i]-'a']--
	}

	for _, v := range alphabet {
		if v != 0 {
			return false
		}
	}
	return true
}

func main() {
	a := ""
	b := ""
	fmt.Println(isAnagram(a, b))
}
