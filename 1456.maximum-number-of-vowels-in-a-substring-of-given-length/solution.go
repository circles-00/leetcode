package main

import (
	"fmt"
	"strings"
)

func isVowel(s string) bool {
	return strings.ToLower(s) == "a" || strings.ToLower(s) == "e" || strings.ToLower(s) == "i" || strings.ToLower(s) == "o" || strings.ToLower(s) == "u"
}

func maxVowels(s string, k int) int {
	left := 0
	right := left + k - 1

	currVowels := 0
	maxVowels := 0

	for i := left; i <= right; i++ {
		if isVowel(string(s[i])) {
			currVowels++
		}
	}

	maxVowels = currVowels

	for right < len(s)-1 {
		if isVowel(string(s[left])) {
			currVowels--
		}

		left++
		right++

		if isVowel(string(s[right])) {
			currVowels++
		}

		if currVowels > maxVowels {
			maxVowels = currVowels
		}
	}

	return maxVowels
}

func main() {
	s := "abciiidef"
	k := 3

	fmt.Println(maxVowels(s, k))
}
