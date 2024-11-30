package tests

import (
	"task2/functions"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomSlice(t *testing.T) {
	testCases := []struct {
		name          string
		size          int
		min           int
		max           int
		expectedError bool
	}{
		{

			name:          "Correct data",
			size:          10,
			min:           -100,
			max:           100,
			expectedError: false,
		},
		{
			name:          "Switch min max",
			size:          5,
			min:           100,
			max:           -100,
			expectedError: false,
		},
		{
			name:          "Error by size",
			size:          0,
			min:           1,
			max:           1,
			expectedError: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			slice, err := functions.GenerateRandomSlice(testCase.size, testCase.min, testCase.max)
			if testCase.expectedError {
				assert.Error(t, err, "expected an error")
				if err == nil {
					t.Error("Expected Error: but got nil")
				}
			} else {
				assert.NoError(t, err)
				assert.Equal(t, len(slice), testCase.size)
			}
		})

	}
}

func TestGetEvenNumbersOnly(t *testing.T) {
	testCases := []struct {
		name        string
		slice       []int
		resultSlice []int
	}{
		{
			name:        "Normal data",
			slice:       []int{1, 2, 3, 4, 5, 6},
			resultSlice: []int{2, 4, 6},
		},
		{
			name:        "Second slice empty",
			slice:       []int{1, 1, 1, 1, 1},
			resultSlice: []int{},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := functions.GetEvenNumbersOnly(testCase.slice)
			assert.Equal(t, result, testCase.resultSlice)
		})
	}
}

func TestMyAppend(t *testing.T) {
	testCases := []struct {
		name        string
		slice       []int
		addedNum    int
		resultSlice []int
	}{
		{
			name:        "Add in not empty slice",
			slice:       []int{1, 2, 3, 4},
			addedNum:    5,
			resultSlice: []int{1, 2, 3, 4, 5},
		},
		{
			name:        "Add in empty slice",
			slice:       []int{},
			addedNum:    99,
			resultSlice: []int{99},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := functions.MyAppend(testCase.slice, testCase.addedNum)
			assert.Equal(t, testCase.resultSlice, result)
		})
	}
}

func TestRemoveElementByIndex(t *testing.T) {
	testCases := []struct {
		name          string
		slice         []int
		deleteIndex   int
		resultSlice   []int
		expectedError bool
	}{
		{
			name:          "Delete mid",
			slice:         []int{0, 1, 2, 3, 4},
			deleteIndex:   2,
			resultSlice:   []int{0, 1, 3, 4},
			expectedError: false,
		},
		{
			name:          "Delete by not exist index",
			slice:         []int{0, 1, 2, 3, 4},
			deleteIndex:   10,
			resultSlice:   []int{0, 1, 2, 3, 4},
			expectedError: true,
		},
		{
			name:          "Empty Slice",
			slice:         []int{},
			deleteIndex:   0,
			resultSlice:   []int{},
			expectedError: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result, err := functions.RemoveElementByIndex(testCase.slice, testCase.deleteIndex)
			if testCase.expectedError {
				assert.Error(t, err, "expected an error")
				if err == nil {
					t.Error("Expected Error, but got nil")
				}
			} else {
				assert.NoError(t, err)
				assert.Equal(t, result, testCase.resultSlice)
			}
		})

	}
}
