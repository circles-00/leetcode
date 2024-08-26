package main

func groupThePeople(groupSizes []int) [][]int {
	finalGroups := [][]int{}
	groups := map[int][]int{}

	for person, groupSize := range groupSizes {
		groups[groupSize] = append(groups[groupSize], person)
	}

	for key, group := range groups {
		count := 0
		groupArr := []int{}
		exceededGroup := []int{}
		for _, person := range group {
			if len(group) > key {
				count++
				exceededGroup = append(exceededGroup, person)
				if count == key {
					count = 0
					finalGroups = append(finalGroups, exceededGroup)
					exceededGroup = []int{}
					continue
				}
			}
			groupArr = append(groupArr, person)
		}

		if len(group) > key {
			continue
		}

		finalGroups = append(finalGroups, groupArr)
	}

	return finalGroups
}

func main() {
	groupSizes := []int{3, 3, 3, 3, 3, 1, 3}

	result := groupThePeople(groupSizes)

	for i, value := range result {
		for _, v := range value {
			println("index:", i, "value:", v)
		}
	}
}
