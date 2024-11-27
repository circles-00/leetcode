package main

import (
	"bytes"
	"fmt"
)

func compressedString(word string) string {
	comp := *bytes.NewBufferString("")

	currPrefix := *bytes.NewBufferString("")

	for i := 0; i < len(word); i++ {
		if len(currPrefix.String()) > 0 && (currPrefix.String()[0] != word[i] || len(currPrefix.String()) == 9) {
			comp.WriteString(fmt.Sprintf("%d%s", len(currPrefix.String()), string(currPrefix.String()[0])))
			currPrefix.Reset()
		}

		currPrefix.WriteString(string(word[i]))
	}

	if len(currPrefix.String()) > 0 {
		comp.WriteString(fmt.Sprintf("%d%s", len(currPrefix.String()), string(currPrefix.String()[0])))
		currPrefix.Reset()
	}

	return comp.String()
}

func main() {
	word := "aaaaaaaaaaaaaabb"
	// word := "abcde"

	fmt.Println(compressedString(word))
}
