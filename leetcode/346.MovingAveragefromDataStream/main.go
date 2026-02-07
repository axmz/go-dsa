package main

import "fmt"

type MovingAverage struct {
	size int
	sum  int
	s    []int
}

func Constructor(size int) MovingAverage {

	return MovingAverage{
		size: size,
		s:    []int{},
	}
}

func (this *MovingAverage) Next(val int) float64 {

	if len(this.s) < this.size {
		this.s = append(this.s, val)
		this.sum += val
	} else {
		this.sum -= this.s[0]
		this.s = this.s[1:]
		this.s = append(this.s, val)
		this.sum += val
	}
	return float64(this.sum) / float64(len(this.s))
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
