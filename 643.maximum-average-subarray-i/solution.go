package main

import "fmt"

func findMaxAverage(nums []int, k int) float64 {
	left := 0
	right := left + k - 1

	currSum := 0
	maxSum := 0

	for i := left; i <= right; i++ {
		currSum += nums[i]
	}

	maxSum = currSum

	for right < len(nums)-1 {
		currSum -= nums[left]
		left++

		right++

		currSum += nums[right]

		if currSum > maxSum {
			maxSum = currSum
		}
	}

	return float64(maxSum) / float64(k)
}

func main() {
	nums := []int{1, 12, -5, -6, 50, 3}
	k := 4

	fmt.Println(findMaxAverage(nums, k))
}
