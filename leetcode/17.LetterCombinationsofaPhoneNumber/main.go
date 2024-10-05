package main

import (
	"fmt"
	"strings"
)

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	m := map[rune]string{
		'2': "abc", '3': "def", '4': "ghi", '5': "jkl",
		'6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz",
	}

	var backtrack func(digits string) []string
	backtrack = func(digits string) []string {
		letters := strings.Split(m[rune(digits[0])], "")
		if len(digits) == 1 {
			return letters
		}

		combinations := []string{}

		for _, i := range letters {
			for _, j := range backtrack(digits[1:]) {
				combinations = append(combinations, i+j)
			}
		}

		return combinations
	}

	return backtrack(digits)
}

func main() {
	fmt.Println(letterCombinations("23"))
}
