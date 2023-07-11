package main

import (
	"fmt"
)

func gcd(a, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}

func gcdOfStrings(str1 string, str2 string) string {
	l1, l2 := len(str1), len(str2)
	s := str1 + str2
	if s != str2+str1 {
		return ""
	}

	g := gcd(l1, l2)

	return s[:g]
}

func main() {
	str1 := "ABCABC"
	str2 := "ABC"
	// str2 := "ABABAB"
	// str1 := "ABAB"
	// str1 := "TAUXXTAUXXTAUXXTAUXXTAUXX"
	// str2 := "TAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXX"

	res := gcdOfStrings(str1, str2)
	fmt.Println(res)
}
