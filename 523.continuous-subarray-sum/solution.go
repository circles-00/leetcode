package main

func checkSubarraySum(nums []int, k int) bool {
	foundRemainders := map[int]int{0: -1}
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		remainder := sum % k

		value, exists := foundRemainders[remainder]

		if !exists {
			foundRemainders[remainder] = i
		}

		if i-value > 1 {
			return true
		}
	}

	return false
}

func main() {
	nums := []int{2, 4, 3}
	k := 6

	println(checkSubarraySum(nums, k))
}
