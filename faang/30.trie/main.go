package main

import "fmt"

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

// Trie can be implemented with array or map
// /////////////////////
// Trie array
// /////////////////////
type Trie struct {
	end      bool // alternatively to end bool we can store the word itself
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
	cur.end = true // there was no need for special case if i == len(word)-1
}

func (this *Trie) Search(word string) bool {
	current := this
	for i := 0; i < len(word); i++ {
		w := word[i] - 'a'
		if t := current.children[w]; t != nil {
			current = t
		} else {
			return false
		}
	}
	return current.end
}

func (this *Trie) StartsWith(prefix string) bool {
	current := this
	for i := 0; i < len(prefix); i++ {
		w := prefix[i] - 'a'
		if t := current.children[w]; t != nil {
			current = t
		} else {
			return false
		}
	}
	return true
}

// /////////////////
// Trie map
// /////////////////
// type Trie struct {
// 	end      bool
// 	children map[byte]*Trie
// }

// func Constructor() Trie {
// 	return Trie{
// 		children: make(map[byte]*Trie),
// 	}
// }

// func (this *Trie) Insert(word string) {
// 	cur := this
// 	for i := 0; i < len(word); i++ {
// 		if t, ok := cur.children[word[i]]; ok {
// 			cur = t
// 		} else {
// 			cur.children[word[i]] = &Trie{
// 				children: make(map[byte]*Trie),
// 			}
// 			cur = cur.children[word[i]]
// 		}
// 	}
// 	cur.end = true // there was no need for special case if i == len(word)-1
// }

// func (this *Trie) Search(word string) bool {
// 	current := this
// 	for i := 0; i < len(word); i++ {
// 		if t, ok := current.children[word[i]]; ok {
// 			current = t
// 		} else {
// 			return false
// 		}
// 	}
// 	return current.end
// }

// func (this *Trie) StartsWith(prefix string) bool {
// 	current := this
// 	for i := 0; i < len(prefix); i++ {
// 		if t, ok := current.children[prefix[i]]; ok {
// 			current = t
// 		} else {
// 			return false
// 		}
// 	}
// 	return true
// }

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
