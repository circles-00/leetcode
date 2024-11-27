package main

import (
	"fmt"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1

	maxArea := 0
	for left < right {
		w := right - left
		h := min(height[left], height[right])

		area := w * h

		if area > maxArea {
			maxArea = area
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}

func main() {
	// height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	height := []int{1, 1}

	fmt.Println(maxArea(height))
}
