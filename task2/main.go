package main

import (
	"fmt"
	"task2/functions"
)

func main() {
	// 1. Создайте слайс целых чисел originalSlice, содержащий 10 произвольных значений, которые генерируются случайным
	randomSlice, err := functions.GenerateRandomSlice(10, -100, 100)
	if err != nil {
		fmt.Println("Error GenerateRandomSlice: %w", err)
		return
	}
	fmt.Println("Слайс случайных чисел: ", randomSlice)

	// 2. напишите функцию sliceExample, которая принимает слайс и возвращает новый слайс, содержащий только четные числа из исходного слайса.
	onlyEventNumbers := functions.GetEvenNumbersOnly(randomSlice)
	fmt.Println("Только чётные числа: ", onlyEventNumbers)

	// 4. Напишите функцию copySlice, которая принимает слайс и возвращает его копию.
	copySlice := functions.CopySlice(onlyEventNumbers)
	fmt.Println("Копия слайса: ", copySlice)

	// 3. Напишите функцию addElements, которая принимает слайс и число
	// 4. Убедитесь, что изменения в оригинальном слайсе не влияют на его копию.
	editSliceBeforeCopy := functions.MyAppend(onlyEventNumbers, -11111)
	editSliceAfterCopy := functions.MyAppend(copySlice, -22222)

	fmt.Println("Изменения в оригинале слайса: ", editSliceBeforeCopy)
	fmt.Println("Изменения в копии слайса: ", editSliceAfterCopy)

	// 5. Напишите функцию removeElement, которая принимает слайс и индекс элемента, который нужно удалить
	removeIndex := len(editSliceAfterCopy) / 2
	sliceAfterRemove, err := functions.RemoveElementByIndex(editSliceAfterCopy, removeIndex)
	if err != nil {
		fmt.Println("Error RemoveElementByIndex: %w", err)
		return
	}
	fmt.Printf("Слайс после удаления index=%d: %v", removeIndex, sliceAfterRemove)
}
