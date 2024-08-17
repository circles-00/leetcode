package main

import (
	"slices"
	"strconv"
	"strings"
)

var words = map[int]string{
	0:   "",
	1:   "One",
	2:   "Two",
	3:   "Three",
	4:   "Four",
	5:   "Five",
	6:   "Six",
	7:   "Seven",
	8:   "Eight",
	9:   "Nine",
	10:  "Ten",
	11:  "Eleven",
	12:  "Twelve",
	13:  "Thirteen",
	14:  "Fourteen",
	15:  "Fifteen",
	16:  "Sixteen",
	17:  "Seventeen",
	18:  "Eighteen",
	19:  "Nineteen",
	20:  "Twenty",
	30:  "Thirty",
	40:  "Forty",
	50:  "Fifty",
	60:  "Sixty",
	70:  "Seventy",
	80:  "Eighty",
	90:  "Ninety",
	100: "One Hundred",
	101: "One Hundred One",
	102: "One Hundred Two",
	103: "One Hundred Three",
	104: "One Hundred Four",
	105: "One Hundred Five",
	106: "One Hundred Six",
	107: "One Hundred Seven",
	108: "One Hundred Eight",
	109: "One Hundred Nine",
}

var additionalStrings = map[int]string{
	0: "",
	1: "Thousand",
	2: "Million",
	3: "Billion",
}

func baseParse(num int) string {
	value, ok := words[num]

	if ok {
		return value
	}

	if num/100 != 0 {
		tens := num/10%10*10 + num%10

		value, ok := words[tens]

		if ok {
			return words[num/100] + " Hundred " + value
		}

		if num%10 != 0 {
			return words[num/100] + " Hundred " + words[num/10%10*10] + " " + words[num%10]
		}

		return words[num/100] + " Hundred " + words[num/10%10*10]
	}

	return words[num/10%10*10] + " " + words[num%10]
}

func reverseString(s string) string {
	reversed := ""
	for i := len(s); i > 0; i-- {
		reversed += string(s[i-1])
	}

	return reversed
}

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}

	if num <= 20 {
		return words[num]
	}

	numberParts := make([]string, 0)

	stringNum := strconv.Itoa(num)

	numPart := ""
	for i := len(stringNum); i > 0; i-- {
		numPart += string(stringNum[i-1])
		if len(numPart) == 3 {
			numberParts = append(numberParts, reverseString(numPart))
			numPart = ""
		}

		if i-1 == 0 && len(numPart) > 0 && len(numPart) < 3 {
			numberParts = append(numberParts, reverseString(numPart))
		}
	}

	slices.Reverse(numberParts)

	word := ""
	for i, part := range numberParts {
		if part == "000" {
			continue
		}
		numberPart, _ := strconv.Atoi(part)
		parsedPart := baseParse(numberPart)
		if len(parsedPart) > 0 && string(parsedPart[len(parsedPart)-1]) == " " {
			parsedPart = parsedPart[:len(parsedPart)-1]
		}

		word += parsedPart + " " + additionalStrings[len(numberParts)-i-1] + " "
	}

	return strings.TrimSpace(word)
}

func main() {
	println(numberToWords(1000000))
}
