package main

import (
	"fmt"
	"main/functions"
)

func main() {

	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}
	result, isCollision := functions.GetCollisions(a, b)
	if isCollision {
		fmt.Println(result)
	}
}
