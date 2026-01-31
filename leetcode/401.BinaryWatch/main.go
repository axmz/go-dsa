package main

import (
	"fmt"
	"math/bits"
)

// complicated.
// without bits.OnesCount it is difficult code to generate combinations
func readBinaryWatch(turnedOn int) []string {
	var res []string

	for h := 0; h < 12; h++ {
		i := bits.OnesCount(uint(h))
		for m := 0; m < 60; m++ {
			j := bits.OnesCount(uint(m))
			if i+j == turnedOn {
				res = append(res, fmt.Sprintf("%d:%02d", h, m))
			}
		}
	}

	return res
}

func main() {
	fmt.Println(readBinaryWatch(1))
}
