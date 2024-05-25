package main

import (
	"fmt"
	"sync"
)

/*
Написать программу, которая конкурентно рассчитает значение квадратов чисел
взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

func main() {
	// Объявляем wg типа WaitGroup(механизм ожидания завершения группы задач)
	var wg sync.WaitGroup

	// Исходный массив
	mas := [5]int{2, 4, 6, 8, 10}

	// Проходим циклом по массиву
	for _, val := range mas {
		// Перед запуском горутины увеличиваем счетчик ожидания на 1
		wg.Add(1)
		// Захват переменной val горутиной(необязательно после версии Golang 1.22)
		go func(val int) {
			// Указываем, что после выполнения горутины счетчик уменьшается на 1
			defer wg.Done()
			fmt.Println(val * val)
		}(val) // Вызов горутины с переменной val
	}
	// Wait() дожидается выполнения всех горутин по счетчику
	wg.Wait()
}
