package main

import "fmt"

type NumArray struct {
	nums       []int
	prefixSums []int
}

func Constructor(nums []int) NumArray {
	prefixSums := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		prefixSums[i+1] = prefixSums[i] + nums[i]
	}
	return NumArray{nums: nums, prefixSums: prefixSums}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.prefixSums[right+1] - this.prefixSums[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
