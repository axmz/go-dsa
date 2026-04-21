package main

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

type NestedIterator struct {
	list []int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{
		list: flatten(nestedList),
	}
}

func flatten(list []*NestedInteger) []int {
	res := []int{}
	for _, n := range list {
		if n.IsInteger() {
			res = append(res, n.GetInteger())
		} else {
			res = append(res, flatten(n.GetList())...)
		}
	}
	return res
}

func (this *NestedIterator) Next() int {
	if this.HasNext() {
		val := this.list[0]
		this.list = this.list[1:]
		return val
	}
	return 0
}

func (this *NestedIterator) HasNext() bool {
	return len(this.list) > 0
}

func main() {

}
