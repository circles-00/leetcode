package main

import "fmt"

func uniqueOccurrences(arr []int) bool {
	occurences := make(map[int]int, 0)

	for _, v := range arr {
		occurences[v]++
	}

	unique := make(map[int]int, 0)
	for _, v := range occurences {
		unique[v]++

		if unique[v] > 1 {
			return false
		}
	}

	return true
}

func main() {
	arr := []int{1, 2, 2, 1, 1, 3}

	fmt.Println(uniqueOccurrences(arr))
}
