package main

import (
	"fmt"
	"sync"
)

/*
Реализовать структуру-счетчик, которая будет инкрементироваться в
конкурентной среде. По завершению программа должна выводить итоговое
значение счетчика.
-----------------------------------------------------------------
Использование мьютекса обеспечит правильную конкруентную запись и исключит data race.
Инкрементиировать - Увеличивать на единицу.
*/

type Counter struct {
	sm    sync.Mutex
	count int
}

func (c *Counter) Increment() {
	c.sm.Lock()
	c.count++
	c.sm.Unlock()
}

func main() {
	var wg sync.WaitGroup
	counter := Counter{}

	// Вызов 100 горутин из цикла
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}
	// Ожидание завершения всех горутин
	wg.Wait()
	// Вывод итогового значения счетчика
	fmt.Printf("Итоговое значение счетчика: %d\n", counter.count)
}
