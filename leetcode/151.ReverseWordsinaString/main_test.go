package main

import (
	"testing"
)

// var s = " a good   example"

// var s = "the sky is blue"

var s = "  hello world  "

func BenchmarkReverseWords(b *testing.B) {
	for n := 0; n < b.N; n++ {
		reverseWords(s)
	}
}

func BenchmarkReverseWordsBuildIn(b *testing.B) {
	for n := 0; n < b.N; n++ {
		reverseWordsBuiltIn(s)
	}
}
