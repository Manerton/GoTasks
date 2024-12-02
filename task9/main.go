package main

import (
	"fmt"
	"main/functions"
	"math/rand"
	"time"
)

func main() {

	countIter := 10

	uintChannel := make(chan uint8)
	floatChannel := make(chan float64)

	go functions.ConvertAndCubing(uintChannel, floatChannel)

	go func() {
		r := rand.New(rand.NewSource(time.Now().Unix()))
		for i := 0; i < countIter; i++ {
			uintChannel <- uint8(r.Intn(256))
		}
		close(uintChannel)
	}()

	for floatNum := range floatChannel {
		fmt.Println(floatNum)
	}

}
