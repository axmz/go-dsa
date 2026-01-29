package main

import "fmt"

type Trie struct {
	end      bool
	children [26]*Trie
}

func NewTrie() Trie {
	return Trie{
		children: [26]*Trie{},
	}
}

func (this *Trie) Insert(word string) {
	cur := this
	for i := 0; i < len(word); i++ {
		w := word[i] - 'a'
		if t := cur.children[w]; t != nil {
			cur = t
		} else {
			cur.children[w] = &Trie{
				children: [26]*Trie{},
			}
			cur = cur.children[w]
		}
	}
	cur.end = true
}

func (this *Trie) Search(word string) bool {
	current := this
	for i := 0; i < len(word); i++ {
		if word[i] == '.' {
			for _, child := range current.children {
				if child != nil && child.Search(word[i+1:]) {
					return true
				}
			}
			return false
		}
		w := word[i] - 'a'
		if t := current.children[w]; t != nil {
			current = t
		} else {
			return false
		}
	}
	return current.end
}

type WordDictionary struct {
	trie *Trie
}

func Constructor() WordDictionary {
	return WordDictionary{
		trie: &Trie{
			children: [26]*Trie{},
		},
	}
}

func (this *WordDictionary) AddWord(word string) {
	this.trie.Insert(word)
}

func (this *WordDictionary) Search(word string) bool {
	return this.trie.Search(word)
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
