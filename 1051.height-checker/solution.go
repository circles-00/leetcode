package main

import "sort"

func heightChecker(heights []int) int {
	expected := make([]int, len(heights))
	copy(expected, heights)

	sort.Ints(heights)

	notInOrder := 0

	for i := 0; i < len(heights); i++ {
		if heights[i] != expected[i] {
			notInOrder += 1
		}
	}

	return notInOrder
}

func main() {
	heights := []int{1, 1, 4, 2, 1, 3}

	println(heightChecker(heights))
}
