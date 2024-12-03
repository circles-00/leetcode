package main

import (
	"fmt"
	"strings"
)

func addSpaces(s string, spaces []int) string {
	var newS strings.Builder

	i, j := 0, 0

	for i < len(s) {
		char := string(s[i])

		if j != len(spaces) && i == spaces[j] {
			newS.WriteString(" ")
			j++
		}

		newS.WriteString(char)
		i++
	}

	return newS.String()
}

func main() {
	s := "LeetcodeHelpsMeLearn"
	spaces := []int{8, 13, 15}

	fmt.Println(addSpaces(s, spaces))
}
