package main

import "fmt"

type children [26]*Trie
type Trie struct {
	value    byte
	end      bool
	children children
	word     string
}

func Constructor() Trie {
	return Trie{}
}

func (trie *Trie) Insert(word string) {
	current := trie
	for i := range word {
		letter := word[i]
		letter_idx := letter - 'a'
		t := current.children[letter_idx]
		if t != nil {
			if i == len(word)-1 {
				t.end = true
				t.word = word
			}
			current = t
		} else {
			newT := &Trie{
				value:    letter,
				end:      false,
				children: children{},
			}
			if i == len(word)-1 {
				newT.end = true
				newT.word = word
			}
			current.children[letter_idx] = newT
			current = newT
		}
	}
}

func (trie *Trie) dfs(word string, res *[]string) {
	if trie == nil {
		return
	}

	if word == "" {
		if trie.end {
			*res = append(*res, trie.word)
			if len(*res) == 3 {
				return
			}
		}

		for i := range trie.children {
			trie.children[i].dfs(word, res)
			if len(*res) == 3 {
				return
			}
		}
	} else {
		trie.children[word[0]-'a'].dfs(word[1:], res)
	}
}

func (trie *Trie) SearchDfs(word string) [][]string {
	res := [][]string{}
	for i := range word {
		subres := make([]string, 0, 3)
		trie.dfs(word[:i+1], &subres)
		res = append(res, subres)
	}
	return res
}

func suggestedProducts(products []string, searchWord string) [][]string {
	trie := new(Trie)
	for _, p := range products {
		trie.Insert(p)
	}

	return trie.SearchDfs(searchWord)
}

func main() {

	products := []string{"mobile", "mouse", "moneypot", "monitor", "mousepad"}
	searchWord := "mouse"

	fmt.Println(suggestedProducts(products, searchWord))
}
