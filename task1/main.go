package main

import (
	"fmt"
	"task1/functionts"
)

func main() {
	// Создает несколько переменных различных типов данных.
	var numDecimal int = 42           // Десятичная система
	var numOctal int = 052            // Восьмеричная система
	var numHexadecimal int = 0x2A     // Шестнадцатиричная система
	var pi float64 = 3.14             // Тип float64
	var name string = "Golang"        // Тип string
	var isActive bool = true          // Тип bool
	var complexNum complex64 = 1 + 2i // Тип complex64
	var unexpected float32 = 334.4444

	// Определяет тип каждой переменной и выводит его на экран.
	functionts.TypeDetection(numOctal)
	functionts.TypeDetection(numHexadecimal)
	functionts.TypeDetection(numDecimal)
	functionts.TypeDetection(pi)
	functionts.TypeDetection(name)
	functionts.TypeDetection(isActive)
	functionts.TypeDetection(complexNum)
	functionts.TypeDetection(unexpected)

	//  Преобразует все переменные в строковый тип и объединяет их в одну строку.
	combined := ""
	combined = functionts.AddToStringAny(combined, numDecimal)
	combined = functionts.AddToStringAny(combined, numOctal)
	combined = functionts.AddToStringAny(combined, numHexadecimal)
	combined = functionts.AddToStringAny(combined, pi)
	combined = functionts.AddToStringAny(combined, isActive)
	combined = functionts.AddToStringAny(combined, complexNum)
	combined = functionts.AddToStringAny(combined, unexpected)

	fmt.Println("combined str: ", combined)

	// Преобразовать эту строку в срез рун.
	sliceRune := functionts.ConvertStringToRuneSlice(combined)
	fmt.Println("slice rune: ", sliceRune)

	// Захэшировать этот срез рун SHA256, добавив в середину соль "go-2024" и вывести результат.
	hashWithSalt, err := functionts.HashSHA256WithSalt(sliceRune, "go-2024")
	if err != nil {
		fmt.Printf("Error HashSHA256WithSalt: %v", err)
	}
	fmt.Println("hash SHA256: ", hashWithSalt)

}
