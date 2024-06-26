package main

import "fmt"

/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее
собственное множество.
-----------------------------------------------------------------
Множество реализовано через мапу с ключом - элемент последовательности строк,
а значение пустая структура, т.к. она не занимает памяти.
*/

func main() {
	// Дано
	array := [5]string{"cat", "cat", "dog", "cat", "tree"}
	// Множество
	set := make(map[string]struct{})
	// Заполняем множество
	for _, value := range array {
		set[value] = struct{}{}
	}
	// Вывод
	for key := range set {
		fmt.Println(key)
	}
}
