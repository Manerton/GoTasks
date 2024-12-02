package main

import (
	"fmt"
	"main/mywaitgroup"
)

func main() {
	wg := mywaitgroup.InitMyWaitGroup()

	countIter := 10

	for i := 0; i < countIter; i++ {
		wg.AddOne()

		go func() {
			defer wg.Done()
			fmt.Println(i)
		}()
	}

	wg.Wait()
}
