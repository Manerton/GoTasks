package main

import (
	"context"
	"fmt"
	"main/functions"
)

func main() {
	countGenerateNum := 10

	context, cancel := context.WithCancel(context.Background())

	randomChan := make(chan int)
	go functions.GenerateRandomNumber(context, randomChan)

	for i := 0; i < countGenerateNum; i++ {
		fmt.Println(<-randomChan)
	}

	cancel()

	value, ok := <-randomChan
	if ok {
		fmt.Println(value)
		fmt.Println("channel was not closed")
	} else {
		fmt.Println(value)
		fmt.Println("Channel is closed!")
	}
}
