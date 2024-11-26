package main

import (
	"fmt"
)

func productExceptSelf(nums []int) []int {
	n := len(nums)
	prefix := make([]int, n)
	suffix := make([]int, n)

	accPrefix := 1
	accSuffix := 1
	for i := 0; i < len(nums); i++ {
		// prefix
		prefix[i] = accPrefix * nums[i]
		accPrefix *= nums[i]

		// suffix
		suffix[len(nums)-1-i] = accSuffix * nums[len(nums)-1-i]
		accSuffix *= nums[len(nums)-1-i]
	}

	product := []int{}

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			product = append(product, suffix[i+1])
			continue
		}

		if i == len(nums)-1 {
			product = append(product, prefix[i-1])
			continue
		}

		product = append(product, prefix[i-1]*suffix[i+1])
	}

	return product
}

func main() {
	nums := []int{1, 2, 3, 4}

	result := productExceptSelf(nums)

	fmt.Println(result)
}
