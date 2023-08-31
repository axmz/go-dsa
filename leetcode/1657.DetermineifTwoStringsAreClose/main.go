package main

import (
	"fmt"
	"sort"
)

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	arr1, arr2 := [26]int{}, [26]int{}
	var bitmap1, bitmap2 int32

	// one for loop
	for i := range word1 {
		arr1[word1[i]-'a']++
		bitmap1 = bitmap1 | (1 << (word1[i] - 'a'))
		arr2[word2[i]-'a']++
		bitmap2 = bitmap2 | (1 << (word2[i] - 'a'))
	}

	// two for loops
	// for _, v := range word1 {
	// 	arr1[v-'a']++
	// 	bitmap1 = bitmap1 | (1 << (v - 'a'))
	// }
	// for _, v := range word2 {
	// 	arr2[v-'a']++
	// 	bitmap2 = bitmap2 | (1 << (v - 'a'))
	// }

	if bitmap1 != bitmap2 {
		return false
	}

	sort.Ints(arr1[:])
	sort.Ints(arr2[:])

	return arr1 == arr2
}

func main() {
	// word1, word2 := "abc", "bca"
	// word1, word2 := "a", "aa"
	word1, word2 := "cabbba", "abbccc"
	res := closeStrings(word1, word2)
	fmt.Println(res)
}
