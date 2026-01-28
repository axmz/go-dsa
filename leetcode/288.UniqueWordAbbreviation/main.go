package main

import "fmt"

type ValidWordAbbr struct {
	v map[string]map[string]struct{}
}

func Constructor(dictionary []string) ValidWordAbbr {

	validator := ValidWordAbbr{
		v: make(map[string]map[string]struct{}),
	}

	for _, w := range dictionary {
		h := validator.hash(w)
		if _, exists := validator.v[h]; !exists {
			validator.v[h] = make(map[string]struct{})
		}
		validator.v[h][w] = struct{}{}
	}

	return validator
}

func (this *ValidWordAbbr) IsUnique(word string) bool {
	h := this.hash(word)
	if _, exists := this.v[h]; !exists {
		return true
	}

	if len(this.v[h]) == 1 {
		_, exists := this.v[h][word]
		return exists
	}

	return false
}

func (this *ValidWordAbbr) hash(word string) string {
	if len(word) <= 2 {
		return word
	}
	return fmt.Sprintf("%c%d%c", word[0], len(word)-2, word[len(word)-1])
}

/**
 * Your ValidWordAbbr object will be instantiated and called as such:
 * obj := Constructor(dictionary);
 * param_1 := obj.IsUnique(word);
 */

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
