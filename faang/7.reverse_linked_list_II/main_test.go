package main

import (
	"testing"
)

var tests = []struct {
	name             string
	linkedListString string
	left             int
	right            int
	result           string
}{
	{
		name:             "1,2,3,4,5,6,7",
		linkedListString: "1,2,3,4,5,6,7",
		left:             2,
		right:            4,
		result:           "1,2,5,4,3,6,7",
	},
}

func TestLengthOfLongestSubstring(t *testing.T) {
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			l := createLinkedListFromString(test.linkedListString)
			rev := reverseBetween(&l, test.left, test.right)
			got := signature(rev)
			if expected := test.result; got != expected {
				t.Fatalf("reverseBetween(%v, %v, %v) returned %v; expected %v", test.linkedListString, test.left, test.right, got, expected)
			}
		})
	}
}
