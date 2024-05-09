package main

import "sort"

func maximumHappinessSum(happiness []int, k int) int64 {
	happinessLevel := 0

	sort.Sort(sort.Reverse(sort.IntSlice(happiness)))

	for i := 0; i < k; i++ {
		calculatedHappiness := happiness[i] - i

		if calculatedHappiness < 0 {
			calculatedHappiness = 0
		}

		happinessLevel += calculatedHappiness
	}

	return int64(happinessLevel)
}

func main() {
	hapiness := []int{
		12, 1, 42,
	}

	k := 3

	println(maximumHappinessSum(hapiness, k))
}
