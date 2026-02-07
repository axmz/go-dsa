package main

import "fmt"

// 10110
// 10111
// 10000
// 10001
func binaryGap(n int) int {
	maxGap := 0
	last := -1
	for i := 0; n > 0; i, n = i+1, n>>1 {
		if n&1 == 1 {
			if last != -1 && i-last > maxGap {
				maxGap = i - last
			}
			last = i
		}
	}
	return maxGap
}

func binaryGap2(n int) int {
	maxGap := 0
	last := -1
	for i := 0; n > 0; i, n = i+1, n>>1 {
		if n&1 == 1 {
			if last == -1 {
				last = i
				continue
			}
			if i-last > maxGap {
				maxGap = i - last
			}
			last = i
		}
	}
	return maxGap
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
