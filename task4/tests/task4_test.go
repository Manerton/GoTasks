package tests

import (
	"main/functions"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDifferenceSliceFor(t *testing.T) {
	testCases := []struct {
		name        string
		firstSlice  []string
		secondSlice []string
		resultSlice []string
	}{
		{
			name:        "One unique",
			firstSlice:  []string{"1", "3", "4", "5"},
			secondSlice: []string{"1", "3", "5"},
			resultSlice: []string{"4"},
		},
		{
			name:        "Multiple unique",
			firstSlice:  []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
			secondSlice: []string{"banana", "date", "fig"},
			resultSlice: []string{"apple", "cherry", "43", "lead", "gno1"},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := functions.DifferenceSliceFor(testCase.firstSlice, testCase.secondSlice)
			assert.Equal(t, result, testCase.resultSlice)
		})

	}
}
