package main

import "fmt"

type TwoSum struct {
	num_counts map[int]int
}

func Constructor() TwoSum {
	return TwoSum{num_counts: make(map[int]int)}
}

func (this *TwoSum) Add(number int) {
	this.num_counts[number]++
}

func (this *TwoSum) Find(value int) bool {
	for num := range this.num_counts {
		complement := value - num
		if complement != num {
			if _, ok := this.num_counts[complement]; ok {
				return true
			}
		} else {
			if this.num_counts[num] > 1 {
				return true
			}
		}
	}
	return false
}

/**
 * Your TwoSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * param_2 := obj.Find(value);
 */

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
