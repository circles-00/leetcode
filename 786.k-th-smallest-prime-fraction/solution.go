package main

import "sort"

func kthSmallestPrimeFraction(arr []int, k int) []int {
	fractions := make([]float64, 0)
	fractionsMap := make(map[float64][]int)

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			floatFraction := float64(arr[i]) / float64(arr[j])
			fractionsMap[floatFraction] = []int{arr[i], arr[j]}
			fractions = append(fractions, floatFraction)
		}
	}

	sort.Float64s(fractions)

	return fractionsMap[fractions[k-1]]
}

func main() {
	//	arr := []int{
	//		1, 2, 3, 5,
	//	}
	arr := []int{1, 7}

	// k := 3
	k := 1

	kthSmallestFraction := kthSmallestPrimeFraction(arr, k)

	println(kthSmallestFraction[0])
	println(kthSmallestFraction[1])
}
