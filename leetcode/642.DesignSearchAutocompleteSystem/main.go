package main

import (
	"fmt"
	"sort"
)

type AutocompleteSystem struct {
	children  map[byte]*AutocompleteSystem // alphabet won't work because we have spaces and special characters
	current   *AutocompleteSystem
	sentences map[string]int
	prefix    []byte
}

func Constructor(sentences []string, times []int) AutocompleteSystem {
	ac := AutocompleteSystem{
		children:  make(map[byte]*AutocompleteSystem),
		sentences: make(map[string]int),
		prefix:    []byte{},
	}
	ac.current = &ac
	for i, sentence := range sentences {
		ac.sentences[sentence] = times[i]
		ac.Insert(sentence, times[i])
	}
	return ac
}

func (this *AutocompleteSystem) Insert(sentence string, time int) {
	current := this
	for i := 0; i < len(sentence); i++ {
		c := sentence[i]
		if t := current.children[c]; t != nil {
			current = t
		} else {
			current.children[c] = &AutocompleteSystem{
				children:  make(map[byte]*AutocompleteSystem),
				sentences: make(map[string]int),
			}
			current = current.children[c]
		}
		current.sentences[sentence] += time
	}
}

func Top3(sentences map[string]int) []string {
	type pair struct {
		sentence string
		count    int
	}

	pairs := make([]pair, 0, len(sentences))
	for s, c := range sentences {
		pairs = append(pairs, pair{s, c})
	}

	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].count != pairs[j].count {
			return pairs[i].count > pairs[j].count
		}
		return pairs[i].sentence < pairs[j].sentence
	})

	result := []string{}
	for i := 0; i < 3 && i < len(pairs); i++ {
		result = append(result, pairs[i].sentence)
	}
	return result
}

func (this *AutocompleteSystem) Input(c byte) []string {
	if c == '#' {
		sentence := string(this.prefix)
		this.Insert(sentence, 1)
		this.current = this
		this.prefix = []byte{}
		return []string{}
	}
	this.prefix = append(this.prefix, c)
	if this.current != nil && this.current.children[c] != nil {
		this.current = this.current.children[c]
		return Top3(this.current.sentences)
	} else {
		this.current = nil
		return []string{}
	}
}

/**
 * Your AutocompleteSystem object will be instantiated and called as such:
 * obj := Constructor(sentences, times);
 * param_1 := obj.Input(c);
 */

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
