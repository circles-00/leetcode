package main

import "fmt"

func moveZeroes(nums []int) {
	left, right := 0, 0

	for right < len(nums) {
		if nums[left] != 0 && nums[right] == 0 {
			left = right
		}

		if nums[right] != 0 && nums[left] == 0 {
			nums[right], nums[left] = nums[left], nums[right]
			left++
		}

		right++
	}
}

func main() {
	// nums := []int{0, 1, 0, 3, 12}
	nums := []int{1, 0, 1}

	moveZeroes(nums)
	fmt.Println(nums)
}
