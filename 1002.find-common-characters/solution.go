package main

import "strings"

func findMaxString(words []string) string {
	maxWord := words[0]

	for i := 1; i < len(words); i++ {
		if len(words[i]) > len(maxWord) {
			maxWord = words[i]
		}
	}

	return maxWord
}

func commonChars(words []string) []string {
	maxString := findMaxString(words)
	commonLetters := make([]string, 0)

	for i := 0; i < len(maxString); i++ {
		isInWord := true
		for j := 0; j < len(words); j++ {
			if !strings.Contains(words[j], string(maxString[i])) {
				isInWord = false
			}
		}

		if isInWord {
			commonLetters = append(commonLetters, string(maxString[i]))
			for j := 0; j < len(words); j++ {
				indexOfChar := strings.Index(words[j], string(maxString[i]))
				words[j] = words[j][:indexOfChar] + words[j][indexOfChar+1:]
			}
		}
	}

	return commonLetters
}

func main() {
	words := []string{"bella", "label", "roller"}

	result := commonChars(words)

	for _, word := range result {
		println(word)
	}
}
