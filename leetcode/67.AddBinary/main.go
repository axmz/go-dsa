package main

import (
	"fmt"
	"math/big"
)

func addBinary2(a string, b string) string {
	A := new(big.Int)
	A.SetString(a, 2)
	B := new(big.Int)
	B.SetString(b, 2)
	C := new(big.Int)
	C = C.Add(A, B)

	return C.Text(2)
}

func addBinary(a string, b string) string {
	A, _ := new(big.Int).SetString(a, 2)
	B, _ := new(big.Int).SetString(b, 2)
	and := new(big.Int).And(A, B)
	xor := new(big.Int).Xor(A, B)
	carry := new(big.Int).Lsh(and, 1)
	zero := big.NewInt(0)

	for carry.Cmp(zero) != 0 {
		and.And(carry, xor)
		xor.Xor(xor, carry)
		carry.Lsh(and, 1)
	}

	return xor.Text(2)
}

func main() {
	a, b := "10101", "10011"
	fmt.Println(addBinary(a, b))
}
