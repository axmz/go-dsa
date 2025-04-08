package main

import "fmt"

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordMap := make(map[string]bool, len(wordList))
	for _, w := range wordList {
		wordMap[w] = true
	}
	if _, ok := wordMap[endWord]; !ok {
		return 0
	}

	q := []string{beginWord}
	level := 1

	getNeighbours := func(word string) []string {
		res := []string{}
		alphabet := 26
		letterA := byte('a')

		for i := 0; i < len(word); i++ {
			for j := 0; j < alphabet; j++ {
				newW := []byte(word)
				newW[i] = letterA + byte(j)
				if ok := wordMap[string(newW)]; ok {
					res = append(res, string(newW))
					wordMap[string(newW)] = false
				}
			}
		}

		return res
	}

	for len(q) > 0 {
		nextLevel := []string{}
		size := len(q)

		for i := 0; i < size; i++ {
			pop := q[0]
			q = q[1:]

			if pop == endWord {
				return level
			}

			neighbours := getNeighbours(pop)
			nextLevel = append(nextLevel, neighbours...)
		}

		level++
		q = nextLevel
	}

	return 0
}

func main() {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}

	fmt.Println(ladderLength(beginWord, endWord, wordList))
}
