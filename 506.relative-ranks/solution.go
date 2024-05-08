package main

import (
	"sort"
	"strconv"
)

func findRelativeRanks(score []int) []string {
	sortedScore := make([]int, len(score))
	copy(sortedScore, score)
	sort.Ints(sortedScore)

	places := make(map[int]int)

	for i := len(sortedScore) - 1; i >= 0; i-- {
		places[sortedScore[i]] = len(sortedScore) - i
	}

	ranks := map[int]string{
		1: "Gold Medal",
		2: "Silver Medal",
		3: "Bronze Medal",
	}

	finalRanking := make([]string, 0)

	for i := 0; i < len(score); i++ {
		relativeRank, found := ranks[places[score[i]]]
		if !found {
			relativeRank = strconv.Itoa(places[score[i]])
		}
		finalRanking = append(finalRanking, relativeRank)
	}

	return finalRanking
}

func main() {
	score := []int{
		10, 3, 8, 9, 4,
	}

	relativeRanks := findRelativeRanks(score)

	for _, rank := range relativeRanks {
		println(rank)
	}
}
