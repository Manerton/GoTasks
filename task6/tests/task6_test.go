package tests

import (
	"context"
	"main/functions"
	"testing"
)

func TestGenerateRandomNumber(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan int)

	go functions.GenerateRandomNumber(ctx, ch)

	preRandNumber := <-ch
	for i := 0; i < 5; i++ {
		randNumber := <-ch
		if randNumber == preRandNumber {
			t.Error("duplicate random number")
		}
		preRandNumber = randNumber
	}

	cancel()
	_, ok := <-ch
	if ok {
		t.Error("channel was not closed")
	}
}
