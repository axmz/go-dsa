package main

import (
	"fmt"
	"godsa/sorting/bubble"
	"godsa/sorting/insertion"
	"godsa/sorting/merge"
	"godsa/sorting/quick"
	"godsa/sorting/selection"
)

func main() {
	arr := []int{1, 4, 6, 2, 6, 7, 19, 22, 11, 1, 0}
	sorted1 := bubble.Sort(arr)
	sorted2 := selection.Sort(arr)
	sorted3 := insertion.Sort(arr)
	sorted4 := insertion.Sort2(arr)
	sorted5 := merge.Sort(arr)
	sorted6 := quick.Sort(arr)
	fmt.Printf("%-12s|  %v \n", "bubble", sorted1)
	fmt.Printf("%-12s|  %v \n", "selection", sorted2)
	fmt.Printf("%-12s|  %v \n", "insertion1", sorted3)
	fmt.Printf("%-12s|  %v \n", "insertion2", sorted4)
	fmt.Printf("%-12s|  %v \n", "merge", sorted5)
	fmt.Printf("%-12s|  %v \n", "quick", sorted6)
	fmt.Printf("%-12s|  %v \n", "unsorted", arr)
}
