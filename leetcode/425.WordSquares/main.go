package main

import "fmt"

type Trie struct {
	words    []string // all words have same length, hence words[]string makes perfect sense
	children [26]*Trie
}

func Constructor() *Trie {
	return &Trie{
		children: [26]*Trie{},
	}
}

func (this *Trie) Insert(word string) {
	cur := this
	cur.words = append(cur.words, word) // add word to root node
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
		cur.words = append(cur.words, word) // add word to each node along the path
	}
}

func (this *Trie) GetWords(prefix []byte) []string {
	if this == nil {
		return nil
	}

	current := this
	for i := 0; i < len(prefix); i++ {
		w := prefix[i] - 'a'
		if t := current.children[w]; t != nil {
			current = t
		} else {
			return nil
		}
	}
	return current.words
}

func wordSquares(words []string) [][]string {
	if len(words) == 0 {
		return nil
	}
	l := len(words[0])

	trie := Constructor()
	for _, word := range words {
		trie.Insert(word)
	}

	res := [][]string{}
	square := make([]string, 0, l)

	var backtrack func(i int)
	backtrack = func(i int) {
		if i == l {
			tmp := make([]string, l)
			copy(tmp, square)
			res = append(res, tmp)
			return
		}

		prefix := make([]byte, i)
		for j := 0; j < i; j++ {
			prefix[j] = square[j][i]
		}

		for _, w := range trie.GetWords(prefix) {
			square = append(square, w)
			backtrack(i + 1)
			square = square[:len(square)-1]
		}
	}

	backtrack(0)
	return res
}

func main() {
	dict := []string{"area", "lead", "wall", "lady", "ball"}
	fmt.Println(wordSquares(dict))
}
