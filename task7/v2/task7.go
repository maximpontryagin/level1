package main

import (
	"fmt"
	"sync"
)

/*
Реализовать конкурентную запись данных в map.
-----------------------------------------------------------------
Способ 2. Использованиe sync.Map. Эффективнее обычной мапы с мьютексом при большом количестве ядер.
*/

func main() {
	var sm sync.Map // объявление переменной типа sync.Map, обеспечивает конкурентную запись.(готовое решение)
	var wg sync.WaitGroup

	// Вызыв горутин, которые будут записывать конкурентно данные в мапу
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(sm *sync.Map, val int) {
			sm.Store(val, val)
			wg.Done()
		}(&sm, i)
	}

	// вызываем функцию Range для созданной мапы. Range принимает функцию для прохождения по мапе.
	wg.Wait() // Ждем выполнения горутины
	sm.Range(func(k, v interface{}) bool {
		fmt.Println("key:", k, ", val:", v)
		return true // if false, Range stops
	})

}
