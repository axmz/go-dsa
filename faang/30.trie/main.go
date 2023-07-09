package main

import "fmt"

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

type children map[rune]*Trie
type Trie struct {
	value    rune // value is not needed
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
	trie := Constructor()
	// trie.Insert("apple")
	// fmt.Println(trie.Search("apple"))   // return True)
	// fmt.Println(trie.Search("app"))     // return False)
	// fmt.Println(trie.StartsWith("app")) // return True)
	// trie.Insert("app")
	// fmt.Println(trie.Search("app"))

	// commands := []string{"insert", "insert", "insert", "insert", "insert", "insert", "search", "search", "search", "search", "search", "search", "search", "search", "search", "startsWith", "startsWith", "startsWith", "startsWith", "startsWith", "startsWith", "startsWith", "startsWith", "startsWith"}
	// args := []string{"app", "apple", "beer", "add", "jam", "rental", "apps", "app", "ad", "applepie", "rest", "jan", "rent", "beer", "jam", "apps", "app", "ad", "applepie", "rest", "jan", "rent", "beer", "jam"}

	trie.Insert("app")
	trie.Insert("apple")
	trie.Insert("beer")
	trie.Insert("add")
	trie.Insert("jam")
	trie.Insert("rental")
	fmt.Println(trie.Search("apps"))
	fmt.Println(trie.Search("app"))
	fmt.Println(trie.Search("ad"))
	fmt.Println(trie.Search("applepie"))
	fmt.Println(trie.Search("rest"))
	fmt.Println(trie.Search("jan"))
	fmt.Println(trie.Search("rent"))
	fmt.Println(trie.Search("beer"))
	fmt.Println(trie.Search("jam"))
	fmt.Println(trie.StartsWith("apps"))
	fmt.Println(trie.StartsWith("app"))
	fmt.Println(trie.StartsWith("ad"))
	fmt.Println(trie.StartsWith("applepie"))
	fmt.Println(trie.StartsWith("rest"))
	fmt.Println(trie.StartsWith("jan"))
	fmt.Println(trie.StartsWith("rent"))
	fmt.Println(trie.StartsWith("beer"))
	fmt.Println(trie.StartsWith("jam"))
}
