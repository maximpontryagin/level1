package main

import "fmt"

/*
Реализовать пересечение двух неупорядоченных множеств.
-----------------------------------------------------------------
В качетсве множеств использованы мапы.
*/

func main() {
	set1 := map[int]int{1: 1, 2: 2, 3: 3, 4: 4, 5: 5} // Множество 1
	set2 := map[int]int{4: 4, 5: 5, 6: 6, 7: 7}       // Множество 2
	intersection := make(map[int]int)                 // Пересечение множеств

	// Поиск перечений set1 и set2 и запись в intersection
	for key1 := range set1 {
		for key2 := range set2 {
			if key1 == key2 {
				intersection[key1] = key1
			}
		}
	}
	// Вывод пересечния множеств
	fmt.Print("Пересечение множеств: ")
	for key := range intersection {
		fmt.Printf("%d ", key)
	}
}
