package main

import "fmt"

type MapSum struct {
	trie [26]*MapSum
	m    map[string]int
	sum  int
}

func Constructor() MapSum {
	return MapSum{
		trie: [26]*MapSum{},
		m:    make(map[string]int),
		sum:  0,
	}
}

func (this *MapSum) Insert(key string, val int) {
	cur := this
	delta := val
	if existingVal, exists := this.m[key]; exists {
		delta -= existingVal
	}
	cur.sum += delta
	for i := 0; i < len(key); i++ {
		w := key[i] - 'a'
		if cur.trie[w] == nil {
			cur.trie[w] = &MapSum{
				trie: [26]*MapSum{},
				sum:  0,
			}
		}
		cur = cur.trie[w]
		cur.sum += delta
	}
	this.m[key] = val
}

func (this *MapSum) Sum(prefix string) int {
	cur := this
	for i := 0; i < len(prefix); i++ {
		w := prefix[i] - 'a'
		if cur.trie[w] == nil {
			return 0
		}
		cur = cur.trie[w]
	}
	return cur.sum
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
