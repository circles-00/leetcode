package main

import "strings"

func uncommonFromSentences(s1 string, s2 string) []string {
	s1Array := strings.Split(s1, " ")
	s2Array := strings.Split(s2, " ")

	words := make(map[string]int)

	for _, word := range s1Array {
		words[word] += 1
	}

	for _, word := range s2Array {
		words[word] += 1
	}

	uncommonWords := make([]string, 0)

	for key, value := range words {
		if value == 1 {
			uncommonWords = append(uncommonWords, key)
		}
	}

	return uncommonWords
}

func main() {
	s1 := "this apple is sweet"
	s2 := "this apple is sour"

	result := uncommonFromSentences(s1, s2)

	for _, word := range result {
		println(word)
	}
}
