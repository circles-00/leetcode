package main

import "fmt"

func pivotIndex(nums []int) int {
	totalSum := 0

	for _, n := range nums {
		totalSum += n
	}

	currSum := 0
	index := -1

	for i := 0; i < len(nums); i++ {
		if totalSum-currSum-nums[i] == currSum {
			index = i
			break
		}

		currSum += nums[i]
	}

	return index
}

func main() {
	// nums := []int{1, 7, 3, 6, 5, 6}
	// nums := []int{1, 2, 3}
	nums := []int{2, 1, -1}

	fmt.Println(pivotIndex(nums))
}
