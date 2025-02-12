package main

import "fmt"

func strStr(haystack string, needle string) int {
	for start := 0; start <= len(haystack)-len(needle); start++ {
		for i := 0; i < len(needle); i++ {
			if haystack[start+i] != needle[i] {
				break
			}
			if i == len(needle)-1 {
				return start
			}
		}
	}

	return -1
}

func main() {
	// haystack := "salbutsadx"
	// needle := "sad"
	// haystack := "mississippi"
	// needle := "issip"
	haystack := "hello"
	needle := "ll"
	// haystack := "aaa"
	// needle := "aaaa"
	// haystack := "a"
	// needle := "a"

	fmt.Println(strStr(haystack, needle))
}
