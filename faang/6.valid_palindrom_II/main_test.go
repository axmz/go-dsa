package main

import (
	"fmt"
	"testing"
)

var tests = []struct {
	input  string
	result bool
}{
	{
		input:  "",
		result: true,
	},
	{
		input:  "a",
		result: true,
	},
	{
		input:  "ab",
		result: true,
	},
	{
		input:  "abc",
		result: false,
	},
	{
		input:  "abca",
		result: true,
	},
	{
		input:  "raceacar",
		result: true,
	},
	{
		input:  "abccdba",
		result: true,
	},
	{
		input:  "abcdefdba",
		result: false,
	},
}

func TestLengthOfLongestSubstring(t *testing.T) {
	for _, test := range tests {
		test := test
		t.Run(test.input, func(t *testing.T) {
			fmt.Println(test)
			t.Parallel()
			len := validPalindrome(test.input)
			// fmt.Printf("LengthOfLongestSubstring(%q) returned %v; expected %v", test.input, len, test.result)
			if got, expected := len, test.result; got != expected {
				t.Fatalf("validPalindrome(%q) returned %v; expected %v", test.input, got, expected)
			}
		})
	}
}
