package main

func removeDuplicates(nums []int) int {
	left := 0

	for right := 0; right < len(nums); right++ {
		if nums[left] != nums[right] {
			left++
			nums[left] = nums[right]
		}
	}

	return len(nums[:left+1])
}

func main() {
	nums := []int{1, 2, 3}

	println(removeDuplicates(nums))
}
