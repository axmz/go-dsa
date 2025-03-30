package main

import "fmt"

func alienOrder(words []string) string {
	adj := make(map[byte][]byte)
	indegree := make(map[byte]int)

	for _, word := range words {
		for _, c := range word {
			indegree[byte(c)] = 0
		}
	}

	for i := 1; i < len(words); i++ {
		prev := words[i-1]
		next := words[i]
		for j := 0; j < len(prev) && j < len(next); j++ {
			if prev[j] == next[j] {
				if j == len(next)-1 && len(prev) > len(next) {
					return ""
				}
				continue
			}
			r1 := prev[j]
			r2 := next[j]
			adj[r1] = append(adj[r1], r2)
			indegree[r2]++

			break
		}
	}

	q := []byte{}
	for k, v := range indegree {
		if v == 0 {
			q = append(q, k)
		}
	}

	s := []byte{}

	for len(q) > 0 {
		pop := q[0]
		q = q[1:]
		s = append(s, pop)
		for _, v := range adj[pop] {
			indegree[v]--
			if indegree[v] == 0 {
				q = append(q, v)
			}
		}

	}

	if len(s) != len(indegree) {
		return ""
	}

	return string(s)
}

func main() {
	words := []string{
		"wxqkj",
		"whqg",
		"cckgh",
		"cdxg",
		"cdxdt",
		"cdht",
		"ktgxt",
		"ktgch",
		"ktdw",
		"ktdc",
		"jqw",
		"jmc",
		"jmg",
	}
	fmt.Println(alienOrder(words))
	words = []string{"abc", "ab"}
	fmt.Println(alienOrder(words))
}
