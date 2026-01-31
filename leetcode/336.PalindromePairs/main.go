package main

import (
	"fmt"
)

type Trie struct {
	idx      int    // index of the word in the original array
	word     string // alternatively to end bool we can store the word itself
	palIdxs  []int
	children [26]*Trie
}

func Constructor() *Trie {
	return &Trie{
		idx:      -1,
		children: [26]*Trie{},
	}
}

func (this *Trie) Insert(word string, idx int) {
	cur := this
	for i := len(word) - 1; i >= 0; i-- {
		if isPalindrome(word[:i+1]) {
			cur.palIdxs = append(cur.palIdxs, idx)
		}
		w := word[i] - 'a'
		if t := cur.children[w]; t != nil {
			cur = t
		} else {
			cur.children[w] = &Trie{idx: -1, children: [26]*Trie{}}
			cur = cur.children[w]
		}
	}
	cur.word = word
	cur.idx = idx
	cur.palIdxs = append(cur.palIdxs, idx)
}

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func palindromePairs(words []string) [][]int {
	var res [][]int
	trie := Constructor()
	for i, word := range words {
		trie.Insert(word, i)
	}

	for i, word := range words {
		cur := trie
		for j := 0; j < len(word); j++ {
			if cur.idx >= 0 && cur.idx != i && isPalindrome(word[j:]) {
				res = append(res, []int{i, cur.idx})
			}
			w := word[j] - 'a'
			if cur.children[w] == nil {
				cur = nil
				break
			}
			cur = cur.children[w]
		}
		if cur == nil {
			continue
		}
		for _, idx := range cur.palIdxs {
			if idx != i {
				res = append(res, []int{i, idx})
			}
		}
	}

	return res
}

func main() {
	strs := []string{"abcd", "dcba", "lls", "s", "sssll"}
	fmt.Println(palindromePairs(strs))
}
