package main

import "fmt"

func hash(s string) string {
	alphabet := [26]int{}
	for i := 0; i < len(s); i++ {
		alphabet[s[i]-'a']++
	}

	return fmt.Sprintf("%v", alphabet)
}

func groupAnagrams(strs []string) [][]string {
	res := [][]string{}
	m := make(map[string][]string)

	for _, s := range strs {
		h := hash(s)
		m[h] = append(m[h], s)
	}

	for _, v := range m {
		res = append(res, v)
	}

	return res
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}
