package main

import (
	"testing"
)

var testsGetLast = map[string]struct {
	input  string
	result string
}{
	"empty string": {
		input:  "",
		result: "",
	},
	"abc#d": {
		input:  "abc#d",
		result: "d",
	},
	"abc#d#": {
		input:  "abc#d#",
		result: "b",
	},
	"abc##": {
		input:  "abc##",
		result: "a",
	},
	"abc###": {
		input:  "abc###",
		result: "",
	},
	"abc####": {
		input:  "abc####",
		result: "",
	},
	"#": {
		input:  "#",
		result: "",
	},
	"##": {
		input:  "##",
		result: "",
	},
}

func TestGetLast(t *testing.T) {
	for name, test := range testsGetLast {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			last, _ := GetLast(test.input)
			if got, expected := last, test.result; got != expected {
				t.Fatalf("GetLast(%q) returned %v; expected %v", test.input, got, expected)
			}
		})
	}
}

var testsBackspaceCompare = map[string]struct {
	input1 string
	input2 string
	result bool
}{
	"empty string": {
		input1: "",
		input2: "",
		result: true,
	},
	"d": {
		input1: "abc#d",
		input2: "abc#d#d",
		result: true,
	},
	"b": {
		input1: "abc#d##",
		input2: "abc#d#",
		result: false,
	},
	"x": {
		input1: "xywrrmp",
		input2: "xywrrm#p",
		result: false,
	},
	"bx": {
		input1: "bbbextm",
		input2: "bbb#extm",
		result: false,
	},
	"c": {
		input1: "a#c",
		input2: "b",
		result: false,
	},
}

func TestBackspaceCompare(t *testing.T) {
	for name, test := range testsBackspaceCompare {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got, expected := BackspaceCompare(test.input1, test.input2), test.result; got != expected {
				t.Fatalf("BackspaceCompare(%q, %q) returned %v; expected %v", test.input1, test.input2, got, expected)
			}
		})
	}
}
