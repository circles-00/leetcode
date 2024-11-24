package main

func strStr(haystack string, needle string) int {
	for i := 0; i < len(haystack); i++ {
		if i > len(haystack)-len(needle) {
			return -1
		}
		if string(haystack[i:i+len(needle)]) == needle {
			return i
		}
	}

	return -1
}

func main() {
	haystack := "sadbutsad"
	needle := "sad"

	// haystack := "leetcode"
	// needle := "leeto"
	println(strStr(haystack, needle))
}
