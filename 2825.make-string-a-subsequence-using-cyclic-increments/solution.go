package main

import "fmt"

func canMakeSubsequence(str1 string, str2 string) bool {
	first, second := 0, 0

	for first < len(str1) {
		if second >= len(str2) {
			return true
		}

		if (str1[first] == str2[second]) || (str1[first]+1 == str2[second]) || (str1[first] == 122 && str2[second] == 97) {
			first++
			second++
			continue
		}

		first++
	}

	return second >= len(str2)
}

func main() {
	str1 := "ozztajlbhjhppdn"
	str2 := "paaajcjhdo"

	fmt.Println(canMakeSubsequence(str1, str2))
}
