package main

import (
	"fmt"
	"math"
	"sort"
)

// func successfulPairs2(spells []int, potions []int, success int64) []int {
// 	res := make([]int, len(spells))
// 	sortedSpells := make([][2]int, len(spells))
// 	for i, v := range spells {
// 		sortedSpells[i] = [2]int{i, v}
// 	}

// 	sort.Slice(sortedSpells, func(i, j int) bool { return sortedSpells[i][1] < sortedSpells[j][1] })
// 	sort.Ints(potions)

// 	count := 0
// 	j := len(potions) - 1

// 	for i := 0; i < len(spells); i++ {
// 		for ; j >= 0; j-- {
// 			if sortedSpells[i][1]*potions[j] >= int(success) {
// 				count++
// 			} else {
// 				break
// 			}
// 		}
// 		res[sortedSpells[i][0]] = count

// 	}

// 	return res
// }

func successfulPairs(spells []int, potions []int, success int64) []int {
	res := make([]int, len(spells))
	sort.Ints(potions)

	for i, v := range spells {
		target := int(math.Ceil(float64(success) / float64(v)))
		index := sort.Search(len(potions), func(i int) bool {
			return potions[i] >= target
		})

		res[i] = len(potions) - index
	}

	return res
}

func main() {
	// spells := []int{5, 1, 3}
	// potions := []int{1, 2, 3, 4, 5}
	// success := 7

	// spells := []int{15, 8, 19}
	// potions := []int{38, 36, 23}
	// success := 328

	spells := []int{3, 1, 2}
	potions := []int{8, 5, 8}
	success := 16

	fmt.Println(successfulPairs(spells, potions, int64(success)))
}
