package functions

func GetCollisions(firstSlice []int, secondSlice []int) ([]int, bool) {
	resultSlice := []int{}
	for _, valFirstSlice := range firstSlice {
		for _, valSecondSlice := range secondSlice {
			if valSecondSlice == valFirstSlice {
				resultSlice = append(resultSlice, valSecondSlice)
			}
		}
	}
	return resultSlice, len(resultSlice) > 0
}
