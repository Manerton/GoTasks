package main

import (
	"fmt"
	"main/mymap"
)

func main() {
	Mymap := mymap.InitMyMap()

	keys := []string{"one", "two", "three"}
	for i := 0; i < len(keys); i++ {
		Mymap.Add(keys[i], i)
	}

	copyMap := Mymap.Copy()
	for _, key := range keys {
		fmt.Println("Exists:", Mymap.Exists(key))
		val, ok := Mymap.Get(key)
		fmt.Println("Get: value:", val, "Ok:", ok)
		Mymap.Remove(key)
		fmt.Println("Exists:", Mymap.Exists(key))
		val, ok = Mymap.Get(key)
		fmt.Println("Get: value:", val, "Ok:", ok)
		val, ok = copyMap.Get(key)
		fmt.Println("Get: value:", val, "Ok:", ok)
	}

}
