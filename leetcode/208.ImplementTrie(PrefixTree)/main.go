package main

type children map[rune]*Trie
type Trie struct {
	value    rune
	end      bool
	children children
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	currentTrie := this
	for i, r := range word {
		if t, ok := currentTrie.children[r]; ok {
			// mark end
			if i == len(word)-1 {
				t.end = true
			}
			currentTrie = t
		} else {
			// mark end
			isEnd := false
			if i == len(word)-1 {
				isEnd = true
			}
			newT := &Trie{value: r, end: isEnd}
			if currentTrie.children == nil {
				currentTrie.children = make(children)
			}
			currentTrie.children[r] = newT
			currentTrie = newT
		}
	}
	// "a"
	// "p"
	// "p"
	// range over rune
	// check if exists in children
	// - if doesn't exist create trie and append to children
	// - if exists do nothing
	// in both cases enter the trie

}

func (this *Trie) Search(word string) bool {
	current := this
	for i, r := range word {
		if t, ok := current.children[r]; ok {
			if i == len(word)-1 && t.end {
				return true
			} else if i == len(word)-1 {
				return false
			} else {
				current = t
			}
		} else {
			return false
		}
	}
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	current := this
	for i, r := range prefix {
		if t, ok := current.children[r]; ok {
			if i == len(prefix)-1 {
				return true
			} else {
				current = t
			}
		} else {
			return false
		}
	}
	return false
}

func main() {

}
