package main

import (
	"main/functions"
	"testing"
)

func generateChannel(data []int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for _, val := range data {
			ch <- val
		}
	}()
	return ch
}

func TestMergeNonEmptyData(t *testing.T) {
	data1 := []int{1, 2, 3}
	data2 := []int{4, 5, 6}
	data3 := []int{7, 8, 9}

	ch1 := generateChannel(data1)
	ch2 := generateChannel(data2)
	ch3 := generateChannel(data3)

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
}

func TestMergeEmptyData(t *testing.T) {
	ch1 := generateChannel([]int{})
	ch2 := generateChannel([]int{})
	ch3 := generateChannel([]int{})

	merged := functions.MergeChannel(ch1, ch2, ch3)

	// Результирующий канал должен быть пустым
	for range merged {
		t.Errorf("expected merged channel to be empty")
	}

}
