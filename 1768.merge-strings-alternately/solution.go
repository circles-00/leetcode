package main

import (
	"bytes"
)

func mergeAlternately(word1 string, word2 string) string {
	biggerWord := word1

	if len(word2) > len(word1) {
		biggerWord = word2
	}

	i := 0

	var sb bytes.Buffer
	sb.WriteString("")

	for i < len(biggerWord) {
		if i >= len(word1) || i >= len(word2) {
			sb.WriteString(string(biggerWord[i:]))
			break
		}
		sb.WriteString(string(word1[i]) + string(word2[i]))
		i++
	}

	return sb.String()
}

func main() {
	word1 := "ab"
	word2 := "pqrs"

	println(mergeAlternately(word1, word2))
}
