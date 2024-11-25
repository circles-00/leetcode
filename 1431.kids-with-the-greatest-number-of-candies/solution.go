package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	kids := []bool{}
	max := candies[0]

	for _, candy := range candies {
		if candy > max {
			max = candy
		}
	}

	for _, candy := range candies {
		kids = append(kids, (candy+extraCandies) >= max)
	}

	return kids
}

func main() {
	candies := []int{2, 3, 5, 1, 3}

	extraCandies := 3

	result := kidsWithCandies(candies, extraCandies)

	for _, r := range result {
		println(r)
	}
}
