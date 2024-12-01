package main

import (
	"fmt"
	"main/functions"
	"math/rand/v2"
)

func main() {

	ch1 := make(chan int)
	go func() {
		defer close(ch1)
		for i := 0; i < 3; i++ {
			ch1 <- rand.Int()
		}
	}()

	ch2 := make(chan int)
	go func() {
		defer close(ch2)
		for i := 0; i < 4; i++ {
			ch2 <- rand.Int()
		}
	}()

	ch3 := make(chan int)
	go func() {
		defer close(ch3)
		for i := 0; i < 5; i++ {
			ch3 <- rand.Int()
		}
	}()

	mergeChannel := functions.MergeChannel(ch1, ch2, ch3)

	for num := range mergeChannel {
		fmt.Println(num)
	}

}
