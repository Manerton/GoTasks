package functions

import (
	"fmt"
	"math/rand/v2"
)

func GenerateRandomSlice(size int, min int, max int) ([]int, error) {
	if size < 1 {
		return nil, fmt.Errorf("size cannot be < 1")
	}
	if min > max {
		min, max = max, min
	}
	result := []int{}
	for i := 0; i < size; i++ {
		result = append(result, rand.IntN(max-min)+min)
	}
	return result, nil
}

func GetEvenNumbersOnly(slice []int) []int {
	result := []int{}
	for _, value := range slice {
		if value%2 == 0 {
			result = append(result, value)
		}
	}
	return result
}

func MyAppend(slice []int, num int) []int {
	return append(slice, num)
}

func CopySlice(slice []int) []int {
	copySlice := make([]int, len(slice))
	copy(copySlice, slice)
	return copySlice
}

func RemoveElementByIndex(slice []int, index int) ([]int, error) {
	if index >= len(slice) || index < 0 {
		return nil, fmt.Errorf("index out of range")
	}
	result := append(slice[:index], slice[index+1:]...)
	return result, nil
}
