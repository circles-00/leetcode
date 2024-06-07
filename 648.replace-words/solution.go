package main

import "strings"

func replaceWords(dictionary []string, sentence string) string {
	newSentence := sentence
	dict := make(map[string]bool)

	for _, entry := range dictionary {
		dict[entry] = true
	}

	prefix := ""
	word := ""
	for index, rune := range sentence {
		char := string(rune)
		if dict[word] && prefix == "" {
			prefix = word
		}

		if char == " " || index == len(sentence)-1 {
			if index == len(sentence)-1 {
				word += string(sentence[index])
			}

			if prefix != "" {
				newSentence = strings.Replace(newSentence, word, prefix, 1)
				prefix = ""
			}

			word = ""
			continue
		}

		word += char

	}

	return newSentence
}

func main() {
	dict := []string{"a", "aa", "aaa", "aaaa"}
	sentence := "a aa a aaaa aaa aaa aaa aaaaaa bbb baba ababa"

	println(replaceWords(dict, sentence))
}
