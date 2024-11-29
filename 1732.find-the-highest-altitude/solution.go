package main

import (
	"fmt"
)

func largestAltitude(gain []int) int {
	sum, max := 0, 0

	for i := 0; i < len(gain); i++ {
		sum += gain[i]
		if sum > max {
			max = sum
		}
	}

	return max
}

func main() {
	gain := []int{-5, 1, 5, 0, -7}

	fmt.Println(largestAltitude(gain))
}
