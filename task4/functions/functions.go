package functions

func DifferenceSliceFor(firstSlice []string, secondSlice []string) []string {
	resultSlice := []string{}
	for i := 0; i < len(firstSlice); i++ {
		isUnique := true
		for j := 0; j < len(secondSlice) && isUnique; j++ {
			if firstSlice[i] == secondSlice[j] {
				isUnique = false
			}
		}
		if isUnique {
			resultSlice = append(resultSlice, firstSlice[i])
		}
	}
	return resultSlice
}

func DifferenceSliceMap(firstSlice []string, secondSlice []string) []string {
	resultSlice := []string{}

	secondSliceMap := map[string]bool{}
	for _, value := range secondSlice {
		secondSliceMap[value] = true
	}

	for _, value := range firstSlice {
		if !secondSliceMap[value] {
			resultSlice = append(resultSlice, value)
		}
	}

	return resultSlice
}
