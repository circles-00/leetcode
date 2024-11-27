package main

import "fmt"

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	if len(s) == 1 && len(t) == 1 {
		return s[0] == t[0]
	}

	left, right := 0, 0

	for right < len(t) && left < len(s) {
		if s[left] == t[right] {
			left++
		}

		right++

	}

	return left == len(s)
}

func main() {
	s := "b"
	t := "abc"

	// s := "abc"
	// t := "ahbgdc"

	// s := "bb"
	// t := "ahbgdc"

	fmt.Println(isSubsequence(s, t))
}
