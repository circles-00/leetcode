package main

import (
	"fmt"
	"strconv"
	"strings"
)

func compress(chars []byte) int {
	if len(chars) == 1 || len(chars) == 0 {
		return len(chars)
	}

	counter := 0
	left, right := 0, 0
	for right < len(chars) {
		if chars[right] != chars[left] {
			if counter == 1 {
				left++
			}

			if counter > 1 {
				cntString := strconv.Itoa(counter)
				for j := 0; j < len(cntString); j++ {
					chars[left+j+1] = cntString[j]
				}
				left += len(cntString) + 1
			}

			chars[left] = chars[right]
			counter = 0
			continue
		}

		right++
		counter++
	}

	if counter == 1 {
		chars[left] = chars[right-1]
	}

	if counter > 1 {
		cntString := strconv.Itoa(counter)
		for j := 0; j < len(cntString); j++ {
			chars[left+j+1] = cntString[j]
		}
		left += len(cntString)
	}

	return len(chars[:left+1])
}

func main() {
	// chars := strings.Join([]string{"a", "a", "b", "b", "c", "c", "c"}, "")
	// chars := strings.Join([]string{"a", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b"}, "")
	// chars := strings.Join([]string{"a", "a", "a", "a", "a", "a", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "c", "c", "c", "c", "c", "c", "c", "c", "c", "c", "c", "c", "c", "c"}, "")
	chars := strings.Join([]string{"o", "o", "o", "o", "o", "o", "o", "o", "o", "o"}, "")
	// chars := strings.Join([]string{"a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a"}, "")

	fmt.Println(compress([]byte(chars)))
}
