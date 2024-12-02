package main

import (
	"bytes"
	"fmt"
)

func isPrefixOfWord(sentence string, searchWord string) int {
	wordId := -1
	currentWord := bytes.NewBufferString("")
	currentWordId := 1

	for i := 0; i < len(sentence); i++ {
		currentWord.WriteString(string(sentence[i]))

		if currentWord.String() == searchWord {
			wordId = currentWordId
			break
		}

		if string(sentence[i]) == " " {
			currentWordId++
			currentWord.Reset()
		}
	}

	return wordId
}

func main() {
	sentence := "i love eating burger"

	searchWord := "burg"

	fmt.Println(isPrefixOfWord(sentence, searchWord))
}
