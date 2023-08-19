package main

import (
	"testing"
)

var roman map[string]int = map[string]int{
	"III":     3,
	"XIV":     14,
	"MCMXCIV": 1994,
	"LVIII":   58,
	"IX":      9,
	"IV":      4,
	"XCIX":    99,
}

func TestRomanToInt(t *testing.T) {
	for k, v := range roman {
		got := romanToInt(k)
		want := v
		if got != want {
			t.Errorf("romanToInt(%v) wanted: %v, got: %v", k, want, got)
		}
	}
}
