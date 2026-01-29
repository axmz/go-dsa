package main

import (
	"fmt"
	"strings"
)

// There is native solution as well using strings.HasPrefix but that would be too easy

type Trie struct {
	end      bool
	children [26]*Trie
}

func Constructor() Trie {
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

func (this *Trie) FindPrefix(word string) int {
	current := this
	for i := 0; i < len(word); i++ {
		w := word[i] - 'a'
		if t := current.children[w]; t != nil {
			if t.end {
				return i
			}
			current = t
		} else {
			return -1
		}
	}
	return -1
}

func replaceWords(dictionary []string, sentence string) string {
	trie := Constructor()
	for _, word := range dictionary {
		trie.Insert(word)
	}

	words := strings.Split(sentence, " ")
	res := make([]string, len(words))
	for i, word := range words {
		if idx := trie.FindPrefix(word); idx != -1 {
			res[i] = word[:idx+1]
		} else {
			res[i] = word
		}
	}
	return strings.Join(res, " ")
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
