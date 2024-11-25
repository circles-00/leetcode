package main

import (
	"bytes"
)

func reverseWords(s string) string {
	words := []string{}

	var currWord bytes.Buffer

	for _, c := range s {
		if string(c) == " " && len(currWord.String()) > 0 {
			words = append(words, currWord.String())
			currWord.Reset()
			continue
		}

		if string(c) != " " {
			currWord.WriteString(string(c))
		}
	}

	if len(currWord.String()) > 0 {
		words = append(words, currWord.String())
	}

	var reversedS bytes.Buffer
	for i := len(words) - 1; i >= 0; i-- {
		reversedS.WriteString(words[i])

		if i != 0 {
			reversedS.WriteString(" ")
		}
	}

	return reversedS.String()
}

func main() {
	s := "a good   example"

	println(reverseWords(s))
}
