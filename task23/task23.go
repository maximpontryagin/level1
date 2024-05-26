package main

import (
	"fmt"
)

/*
Удалить i-ый элемент из слайса.
-----------------------------------------------------------------
Использую срезы для удаления i-ого элемента
*/

func remove(slice []int, i int) []int {
	// Проверка на выход за границы слайса
	if i < 0 || i >= len(slice) {
		return slice
	}
	// Удаляем i-ый элемент. ... - нужно для распоковки элементов среза
	return append(slice[:i], slice[i+1:]...)
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Начальный слайс:", slice)

	// Удаление 2-го элемента (индекс 1)
	indexToRemove := 1
	slice = remove(slice, indexToRemove)
	fmt.Println("Удален второй элемент:", slice)

	// Удаление 4-го элемента (индекс 3)
	indexToRemove = 3
	slice = remove(slice, indexToRemove)
	fmt.Println("Удален четвертный элемент:", slice)
}
