package main

func majorityElement(nums []int) int {
	count := map[int]int{}

	for _, num := range nums {
		count[num] += 1
		if count[num] > len(nums)/2 {
			return num
		}
	}

	return 0
}

func main() {
	nums := []int{3, 2, 3}

	println(majorityElement(nums))
}
