package main

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	s map[int]int
	a []int
	// perhaps seeding rand once here is a potential optimization
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		s: make(map[int]int),
		a: []int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, exists := this.s[val]; exists {
		return false
	}
	this.a = append(this.a, val)
	this.s[val] = len(this.a) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, exists := this.s[val]; !exists {
		return false
	}
	last := this.a[len(this.a)-1]
	idx := this.s[val]
	this.a[idx] = last
	this.s[last] = idx
	this.a = this.a[:len(this.a)-1]
	delete(this.s, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	randIdx := rand.Intn(len(this.a))
	return this.a[randIdx]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
