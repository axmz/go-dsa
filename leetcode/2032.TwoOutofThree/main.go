package main

import (
	"fmt"
	"math/big"
)

// used big.Int because there is no uint128
// arr[101] is prolly better than big.Int

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	bitset1 := big.Int{}
	bitset2 := big.Int{}
	bitset3 := big.Int{}

	for _, num := range nums1 {
		bitset1.SetBit(&bitset1, num, 1)
	}
	for _, num := range nums2 {
		bitset2.SetBit(&bitset2, num, 1)
	}
	for _, num := range nums3 {
		bitset3.SetBit(&bitset3, num, 1)
	}

	resultSet := big.Int{}
	temp := big.Int{}
	resultSet.Or(&resultSet, temp.And(&bitset1, &bitset2))
	resultSet.Or(&resultSet, temp.And(&bitset1, &bitset3))
	resultSet.Or(&resultSet, temp.And(&bitset2, &bitset3))

	result := []int{}
	for i := 1; i <= 100; i++ {
		if resultSet.Bit(i) == 1 {
			result = append(result, i)
		}
	}
	return result
}

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
