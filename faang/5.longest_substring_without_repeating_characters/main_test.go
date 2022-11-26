package main

import (
	"fmt"
	"testing"
)

var testsLongestSubstring = []struct {
	input  string
	result int
}{
	{
		input:  "",
		result: 0,
	},
	{
		input:  "bbbbb",
		result: 1,
	},
	{
		input:  "pwwkew",
		result: 3,
	},
	{
		input:  "abcbdaac",
		result: 4,
	},
	{
		input:  "bbtablud",
		result: 6,
	},
	{
		input:  "abcabcbb",
		result: 3,
	},
}

func TestLengthOfLongestSubstring(t *testing.T) {
	for _, test := range testsLongestSubstring {
		test := test
		t.Run(test.input, func(t *testing.T) {
			fmt.Println(test)
			t.Parallel()
			len := LengthOfLongestSubstring(test.input)
			// fmt.Printf("LengthOfLongestSubstring(%q) returned %v; expected %v", test.input, len, test.result)
			if got, expected := len, test.result; got != expected {
				t.Fatalf("LengthOfLongestSubstring(%q) returned %v; expected %v", test.input, got, expected)
			}
		})
	}
}
