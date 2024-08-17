package main

const (
	maxInteger = 2147483647
	minInteger = -2147483648
)

func myAtoi(s string) int {
	digitsMap := map[int]int{
		48: 0,
		49: 1,
		50: 2,
		51: 3,
		52: 4,
		53: 5,
		54: 6,
		55: 7,
		56: 8,
		57: 9,
	}

	number := ""
	sign := "+"
	signsRead := 0
	for i, rune := range s {
		// Exit if encountered a non-digit before reading digits
		if len(number) == 0 && (rune < 48 || rune > 57) && rune != 32 && rune != 45 && rune != 43 {
			break
		}
		if rune == 45 || rune == 43 {
			signsRead += 1
		}

		if signsRead > 1 || (signsRead == 1 && rune == 32) {
			break
		}

		// Exit if encountered a non-digit after reading some digits
		if len(number) != 0 && (rune < 48 || rune > 57) {
			break
		}

		if rune >= 48 && rune <= 57 {
			number += string(rune)
			if len(number) == 1 && i != 0 && (string(s[i-1]) == "-" || string(s[i-1]) == "+") {
				sign = string(s[i-1])
			}
		}
	}

	if len(number) == 0 {
		return 0
	}

	integer := 0

	for _, digit := range number {
		integer = integer*10 + digitsMap[int(digit)]
		if sign == "+" && integer > maxInteger {
			return maxInteger
		}

		if sign == "-" && integer*-1 < minInteger {
			return minInteger
		}
	}

	if sign == "-" {
		integer *= -1
	}

	return integer
}

func main() {
	s := "  +  413"

	println(myAtoi(s))
}
