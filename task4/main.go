package main

import (
	"fmt"
	"main/functions"
	"main/generateWords"
	"time"
)

func main() {

	// slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	// slice2 := []string{"banana", "date", "fig"}

	slice1 := generateWords.GenerateWords(10000)
	slice2 := generateWords.GenerateWords(5000)

	start := time.Now()
	result := functions.DifferenceSliceFor(slice1, slice2)
	duration := time.Since(start)
	fmt.Println(result, "Time:", duration.Nanoseconds(), "Nanoseconds")

	start = time.Now()
	result = functions.DifferenceSliceMap(slice1, slice2)
	duration = time.Since(start)
	fmt.Println(result, "Time:", duration.Nanoseconds(), "Nanoseconds")

}
