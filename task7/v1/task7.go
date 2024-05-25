package main

import (
	"fmt"
	"sync"
)

/*
Реализовать конкурентную запись данных в map.
-----------------------------------------------------------------
Способ 1. Использование мьютекса
*/

type SafeMap struct {
	mu sync.Mutex
	m  map[int]int
}

func (sm *SafeMap) Set(key int, value int) {
	sm.mu.Lock()         // Блокировка мьютекса для обеспечения уникального доступа к горутине
	defer sm.mu.Unlock() // В конце разблокируем мьютекс
	sm.m[key] = value    // Записываем значение
}

func main() {
	sm := SafeMap{m: make(map[int]int)}
	var wg sync.WaitGroup

	// Вызыв горутин, которые будут записывать конкурентно данные в мапу
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			sm.Set(i, i) // Добавляем конкурентно ключ - значение
		}(i)
	}

	wg.Wait() // Ждем выполнения горутины
	fmt.Print(sm.m)
}
