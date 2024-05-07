package main

import "slices"

func findMaxK(nums []int) int {
	k := -1

	for _, num := range nums {
		if slices.Contains(nums, num*-1) && num > k {
			k = num
		}
	}

	return k
}

func main() {
	// nums := []int{-1, 2, -3, 3}
	// nums := []int{-1, 10, 6, 7, -7, 1}
	nums := []int{-10, 8, 6, 7, -2, -3}

	println(findMaxK(nums))
}
