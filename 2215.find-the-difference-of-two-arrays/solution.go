package main

import "fmt"

func findDifference(nums1 []int, nums2 []int) [][]int {
	n1 := make(map[int]int, 0)
	n2 := make(map[int]int, 0)

	matrix := make([][]int, 2)

	for _, n := range nums1 {
		if n1[n] == 0 {
			n1[n]++
		}
	}

	for _, n := range nums2 {
		if n2[n] == 0 {
			n2[n]++
		}
	}

	for _, n := range nums1 {
		if n2[n] == 0 && n1[n] == 1 {
			matrix[0] = append(matrix[0], n)
			n1[n]--
		}
	}

	for _, n := range nums2 {
		if n1[n] == 0 && n2[n] == 1 {
			matrix[1] = append(matrix[1], n)
			n2[n]--
		}
	}

	return matrix
}

func main() {
	nums1 := []int{1, 2, 3, 3}
	nums2 := []int{1, 1, 2, 2}

	fmt.Println(findDifference(nums1, nums2))
}
