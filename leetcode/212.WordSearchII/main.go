package main

import "fmt"

type Trie struct {
	word     string
	children [26]*Trie
}

func Constructor() *Trie {
	return &Trie{
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
	cur.word = word // alternatively to end bool, we can store the word itself
}

var directions = [][]int{
	{0, 1},  // right
	{1, 0},  // down
	{0, -1}, // left
	{-1, 0}, // up
}

func findWords(board [][]byte, words []string) []string {
	trie := Constructor()
	for _, word := range words {
		trie.Insert(word)
	}

	rows := len(board)
	cols := len(board[0])
	// whenever there is visited consider marking in place
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	var res []string
	// be careful where you define the func.
	// make sure it is not in the loop
	var search func(r, c int, trie *Trie)
	search = func(r, c int, trie *Trie) {
		letter := board[r][c]
		idx := letter - 'a'
		child := trie.children[idx] // access dirrectly, no need to loop over all children
		if child != nil {
			visited[r][c] = true
			if child.word != "" {
				res = append(res, child.word)
				child.word = ""
			}
			for _, dir := range directions {
				newR := r + dir[0]
				newC := c + dir[1]
				if newR >= 0 && newR < rows && newC >= 0 && newC < cols && !visited[newR][newC] {
					search(newR, newC, child)
				}
			}
			visited[r][c] = false
		}
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			search(r, c, trie)
		}
	}

	return res
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
