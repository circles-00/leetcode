package main

func lemonadeChange(bills []int) bool {
	const LEMONADE_COST int = 5
	availableChange := make(map[int]int)

	for _, bill := range bills {
		availableChange[bill] += 1

		if bill == LEMONADE_COST {
			continue
		}

		if bill == 10 {
			if availableChange[5] >= 1 {
				availableChange[5] -= 1
				continue
			}
		}

		if bill == 20 {
			if availableChange[10] >= 1 && availableChange[5] >= 1 {
				availableChange[5] -= 1
				availableChange[10] -= 1

				continue
			}

			if availableChange[5] >= 3 {
				availableChange[5] -= 3
				continue
			}
		}

		return false
	}

	return true
}

func main() {
	// bills := []int{5, 5, 10, 10, 20}
	// bills := []int{5, 5, 5, 10, 20}
	bills := []int{10, 10}

	println(lemonadeChange(bills))
}
