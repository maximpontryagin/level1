package main

import "fmt"

/*
Реализовать быструю сортировку массива (quicksort) встроенными методами
языка.
-----------------------------------------------------------------
*/

// BinarySearch выполняет бинарный поиск в отсортированном массиве.
// Возвращает индекс целевого элемента, если он найден, иначе возвращает -1.
func BinarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		// Проверяем, находится ли целевой элемент в середине
		if arr[mid] == target {
			return mid
		}

		// Если целевой элемент больше, игнорируем левую половину и ищим в правой части
		if arr[mid] < target {
			left = mid + 1
		} else {
			// Если целевой элемент меньше, игнорируем правую половину и ищем в левой части
			right = mid - 1
		}
	}

	// Целевой элемент не найден в массиве
	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 7

	result := BinarySearch(arr, target)

	if result != -1 {
		fmt.Printf("Элемент %d найден на индексе %d\n", target, result)
	} else {
		fmt.Printf("Элемент %d не найден в массиве\n", target)
	}
}
