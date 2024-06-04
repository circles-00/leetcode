package main

func longestPalindrome(s string) int {
	if len(s) == 1 {
		return len(s)
	}

	frequencies := make(map[rune]int)

	evenPairs := 0

	for _, rune := range s {
		frequencies[rune] += 1
		if frequencies[rune]%2 == 0 {
			evenPairs += 1
		}
	}

	if evenPairs*2 == len(s) {
		return len(s)
	}

	return evenPairs*2 + 1
}

func main() {
	s := "AAAZZZ"

	println(longestPalindrome(s))
}
