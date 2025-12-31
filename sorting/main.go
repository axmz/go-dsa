package main

import (
	"fmt"
	"godsa/sorting/bubble"
	"godsa/sorting/bucket"
	"godsa/sorting/counting"
	"godsa/sorting/heap"
	"godsa/sorting/insertion"
	"godsa/sorting/merge"
	"godsa/sorting/quick"
	"godsa/sorting/radix"
	"godsa/sorting/selection"
)

func main() {
	arr := []int{1, 4, 6, 201, 36, 7, 1019, 22, 11, 21, 0}
	sorted1 := bubble.Sort(arr)
	sorted2 := selection.Sort(arr)
	sorted3 := insertion.Sort(arr)
	sorted4 := radix.Sort(arr)
	sorted5 := merge.Sort(arr)
	sorted6 := quick.Sort(arr)
	sorted7 := heap.Sort(arr)
	sorted8 := counting.Sort(arr)
	sorted9 := bucket.Sort(arr)
	fmt.Printf("%-12s|  %v \n", "unsorted", arr)
	fmt.Printf("%-12s|  %v \n", "bubble", sorted1)
	fmt.Printf("%-12s|  %v \n", "selection", sorted2)
	fmt.Printf("%-12s|  %v \n", "insertion", sorted3)
	fmt.Printf("%-12s|  %v \n", "merge", sorted5)
	fmt.Printf("%-12s|  %v \n", "quick", sorted6)
	fmt.Printf("%-12s|  %v \n", "heap", sorted7)
	fmt.Printf("%-12s|  %v \n", "counting", sorted8)
	fmt.Printf("%-12s|  %v \n", "bucket", sorted9)
	fmt.Printf("%-12s|  %v \n", "radix", sorted4)
}
