package main

func reversePrefix(word string, ch byte) string {
	index := -1
	wordWithoutPrefix := ""
	for idx, char := range word {
		if string(char) == string(ch) && index == -1 {
			index = idx
			continue
		}

		if index != -1 {
			wordWithoutPrefix += string(char)
		}
	}

	newPrefix := ""
	for i := index; i >= 0; i-- {
		newPrefix += string(word[i])
	}

	if index == -1 {
		return word
	}

	return newPrefix + wordWithoutPrefix
}

func main() {
	word := "abcdefd"
	var ch byte = 'd'

	println(reversePrefix(word, ch))
}
