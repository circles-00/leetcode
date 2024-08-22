package main

import (
	"math/bits"
)

func findComplement(num int) int {
	bitLength := bits.Len(uint(num))
	mask := (1 << bitLength) - 1

	return num ^ mask
}

func main() {
	num := 5

	println(findComplement(num))
}
