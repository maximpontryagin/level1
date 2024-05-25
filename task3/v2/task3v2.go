package main

import (
	"fmt"
	"sync"
)

/*
Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
квадратов(2^2+3^2+4^2….) с использованием конкурентных вычислений.

-------------------------------------------------------------------
Использование каналов. Каналы обеспечивают безопасную передачу данных,
синхронизацию горутин и предотвращают состояние гонки
*/

func main() {
	var wg sync.WaitGroup
	// Переменная для накопления суммы
	sum := 0
	ch := make(chan int)
	// Массив последовательности чисел из дано
	mas := [5]int{2, 4, 6, 8, 10}
	for _, val := range mas {
		// Перед запуском горутины увеличиваем счетчик ожидания на 1
		wg.Add(1)
		go func(val int) {
			// Указываем, что после выполнения горутины счетчик уменьшается на 1
			defer wg.Done()
			ch <- val * val
		}(val) // Вызов горутины с переменной val
	}

	// Закрытие канала после завершения всех горутин для предотвращения deadlock
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Когда закроется канал произойдет выход из цикла
	for value := range ch {
		sum += value
	}

	// Wait() дожидается выполнени я всех горутин по счетчику
	wg.Wait()
	fmt.Println(sum) // sum = 220
}