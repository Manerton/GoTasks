package main

import (
	"main/functions"
	"math/rand"
	"testing"
	"time"
)

func generateTestChannel(data []int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for _, val := range data {
			ch <- val
		}
	}()
	return ch
}

func TestMergeChannel(t *testing.T) {
	rand.NewSource(time.Now().UnixNano())

	t.Run("Merge multiple non-empty channels", func(t *testing.T) {
		data1 := []int{1, 2, 3}
		data2 := []int{4, 5, 6}
		data3 := []int{7, 8, 9}

		ch1 := generateTestChannel(data1)
		ch2 := generateTestChannel(data2)
		ch3 := generateTestChannel(data3)

		merged := functions.MergeChannel(ch1, ch2, ch3)

		expected := map[int]bool{}
		for _, val := range append(append(data1, data2...), data3...) {
			expected[val] = true
		}

		// Проверяем, что все элементы из исходных каналов есть в результирующем
		for val := range merged {
			if !expected[val] {
				t.Errorf("unexpected value %d in merged channel", val)
			}
			delete(expected, val)
		}

		if len(expected) > 0 {
			t.Errorf("not all expected values were found in merged channel: %v", expected)
		}
	})

	t.Run("Merge with empty channels", func(t *testing.T) {
		data1 := []int{10, 20}
		data2 := []int{}
		data3 := []int{30}

		ch1 := generateTestChannel(data1)
		ch2 := generateTestChannel(data2)
		ch3 := generateTestChannel(data3)

		merged := functions.MergeChannel(ch1, ch2, ch3)

		expected := map[int]bool{
			10: true,
			20: true,
			30: true,
		}

		// Проверяем, что значения из непустых каналов есть в результирующем
		for val := range merged {
			if !expected[val] {
				t.Errorf("unexpected value %d in merged channel", val)
			}
			delete(expected, val)
		}

		if len(expected) > 0 {
			t.Errorf("not all expected values were found in merged channel: %v", expected)
		}
	})

	t.Run("Merge with all empty channels", func(t *testing.T) {
		ch1 := generateTestChannel([]int{})
		ch2 := generateTestChannel([]int{})
		ch3 := generateTestChannel([]int{})

		merged := functions.MergeChannel(ch1, ch2, ch3)

		// Результирующий канал должен быть пустым
		for range merged {
			t.Errorf("expected merged channel to be empty")
		}
	})
}
