package main

func threeConsecutiveOdds(arr []int) bool {
	odds := 0
	for _, num := range arr {
		if num%2 == 0 {
			odds = 0
			continue
		}

		odds += 1
		if odds >= 3 {
			return true
		}
	}

	return false
}

func main() {
	arr := []int{2, 6, 4, 1}

	println(threeConsecutiveOdds(arr))
}
