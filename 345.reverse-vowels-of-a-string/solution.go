package main

import (
	"bytes"
	"strings"
)

func isVowel(s string) bool {
	return strings.ToLower(s) == "a" || strings.ToLower(s) == "e" || strings.ToLower(s) == "i" || strings.ToLower(s) == "o" || strings.ToLower(s) == "u"
}

func reverseVowels(s string) string {
	left := 0
	right := len(s) - 1
	replacements := make(map[int]string, 0)

	for left < right {
		if isVowel(string(s[left])) && isVowel(string(s[right])) {
			replacements[left] = string(s[right])
			replacements[right] = string(s[left])
			left++
			right--
			continue
		}

		if !isVowel(string(s[left])) {
			left++
		}

		if !isVowel(string(s[right])) {
			right--
		}
	}

	var str bytes.Buffer

	for i, c := range s {
		r, ok := replacements[i]
		if ok {
			str.WriteString(r)
			continue
		}

		str.WriteString(string(c))
	}

	return str.String()
}

func main() {
	s := "IceCreAm"

	println(reverseVowels(s))
}
