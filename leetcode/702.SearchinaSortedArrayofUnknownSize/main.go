package main

import "fmt"

/**
 * // This is the ArrayReader's API interface.
 * // You should not implement it, or speculate about its implementation
 * type ArrayReader struct {
 * }
 *
 * func (this *ArrayReader) get(index int) int {}
 */

func search(reader ArrayReader, target int) int {
	l, r := 0, 1
	max := (1 << 31) - 1
	for {
		data := reader.get(r)
		if l == r && data != target {
			return -1
		}
		switch {
		case data == target:
			return r
		case data < target:
			l = r
			r <<= 1
		case data > target, data == max:
			r = l + (r-l)/2
		}
	}
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
