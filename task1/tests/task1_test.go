package tests

import (
	"fmt"
	"task1/functionts"
	"testing"
)

func TestTypeDetection(t *testing.T) {
	testCases := []struct {
		object   interface{}
		expected string
	}{
		{42, "int"},
		{float64(3.14), "float64"},
		{float32(3.15), "float32"},
		{true, "bool"},
		{"hello", "string"},
		{complex64(1 + 2i), "complex64"},
	}

	for _, testCase := range testCases {
		resultType := functionts.TypeDetection(testCase.object)
		if resultType != testCase.expected {
			t.Errorf("expected %s but got %s", testCase.expected, resultType)
		}
	}
}

func TestAddToStringAny(t *testing.T) {
	testCases := []struct {
		name      string
		sliceObj  []any
		resultStr string
	}{
		{
			name:      "with string/digit",
			sliceObj:  []any{"14", "Hello", 24, 55, "Hel2lo", 23.54},
			resultStr: "14Hello2455Hel2lo23.54",
		},
		{
			name:      "with boolean",
			sliceObj:  []any{complex64(1 + 3i), "Go", 24.245, false, "-1001-", true},
			resultStr: "(1+3i)Go24.245false-1001-true",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := ""
			for _, obj := range testCase.sliceObj {
				result = functionts.AddToStringAny(result, obj)
			}
			if result != testCase.resultStr {
				t.Errorf("expected %s but got %s", testCase.resultStr, result)
			}
		})

	}
}

func TestConvertStringToRuneSlice(t *testing.T) {
	testCases := []struct {
		name       string
		str        string
		resultRune []rune
	}{
		{
			name:       "Only eng symbol",
			str:        "HelloWorld",
			resultRune: []rune("HelloWorld"),
		},
		{
			name:       "Mix",
			str:        "Go2024БуквыТЕСТ",
			resultRune: []rune("Go2024БуквыТЕСТ"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := functionts.ConvertStringToRuneSlice(testCase.str)
			for i, r := range result {
				if r != testCase.resultRune[i] {
					t.Errorf("Expected %c, got %c", testCase.resultRune[i], r)
				}
			}
		})

	}
}

func TestHashWithSalt(t *testing.T) {
	testCases := []struct {
		name          string
		str           string
		salt          string
		expectedError bool
	}{
		{
			name:          "Without Error",
			str:           "notError",
			salt:          "go-2024",
			expectedError: false,
		},
		{
			name:          "With Error",
			str:           "Error",
			salt:          "",
			expectedError: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			runes := functionts.ConvertStringToRuneSlice(testCase.str)
			hash, err := functionts.HashSHA256WithSalt(runes, testCase.salt)
			if testCase.expectedError {
				fmt.Println("Expected err:", err)
			} else {
				expectedLength := 32
				if len(hash) != expectedLength {
					t.Errorf("Expected hash of length %d, got %d", expectedLength, len(hash))
				}
			}
		})

	}
}
