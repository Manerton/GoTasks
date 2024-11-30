package functions

import (
	"main/functions"
	"reflect"
	"testing"
)

func TestGetCollisions(t *testing.T) {
	testCases := []struct {
		name           string
		firstSlice     []int
		secondSlice    []int
		expectedResult []int
		expectedBool   bool
	}{
		{
			name:           "No collisions",
			firstSlice:     []int{1, 2, 3},
			secondSlice:    []int{4, 5, 6},
			expectedResult: []int{},
			expectedBool:   false,
		},
		{
			name:           "Multiple collisions",
			firstSlice:     []int{1, 2, 3, 4},
			secondSlice:    []int{3, 4, 5, 6},
			expectedResult: []int{3, 4},
			expectedBool:   true,
		},
		{
			name:           "Empty slices",
			firstSlice:     []int{},
			secondSlice:    []int{},
			expectedResult: []int{},
			expectedBool:   false,
		},
		{
			name:           "Duplicate values in slices",
			firstSlice:     []int{1, 1, 2, 3},
			secondSlice:    []int{1, 3, 3, 4},
			expectedResult: []int{1, 1, 3, 3},
			expectedBool:   true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result, hasCollisions := functions.GetCollisions(testCase.firstSlice, testCase.secondSlice)
			if !reflect.DeepEqual(result, testCase.expectedResult) {
				t.Errorf("expected result %v, got %v", testCase.expectedResult, result)
			}
			if hasCollisions != testCase.expectedBool {
				t.Errorf("expected bool %v, got %v", testCase.expectedBool, hasCollisions)
			}
		})
	}
}
