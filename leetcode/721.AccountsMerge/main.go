package main

import (
	"fmt"
	. "godsa/utils/disjointset"
	"sort"
)

func accountsMerge(accounts [][]string) [][]string {
	n := len(accounts)
	u := NewUnion(n)

	emailToID := make(map[string]int)
	for i, account := range accounts {
		for _, email := range account[1:] {
			if id, ok := emailToID[email]; !ok {
				emailToID[email] = i
			} else {
				u.Union(i, id)
			}
		}
	}

	idToEmails := make(map[int][]string)
	for email, id := range emailToID {
		root := u.Find(id)
		idToEmails[root] = append(idToEmails[root], email)
	}

	res := make([][]string, 0)
	for id, emails := range idToEmails {
		account := []string{accounts[id][0]}
		account = append(account, emails...)
		sort.Strings(account[1:])
		res = append(res, account)
	}
	return res
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
