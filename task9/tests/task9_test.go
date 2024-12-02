package tests

import (
	"main/functions"
	"math"
	"testing"
)

func TestConvertAndCubing(t *testing.T) {
	startData := []uint8{2, 3, 4, 5, 6}
	cubingData := map[float64]bool{}
	for _, val := range startData {
		num := float64(val)
		cubingData[math.Pow(num, 3)] = true
	}

	uintChannel := make(chan uint8)
	floatChannel := make(chan float64)

	go functions.ConvertAndCubing(uintChannel, floatChannel)

	go func() {
		for _, val := range startData {
			uintChannel <- val
		}
		close(uintChannel)
	}()

	for cubNum := range floatChannel {
		if !cubingData[cubNum] {
			t.Error("Cubing is not work")
		}
	}
}
