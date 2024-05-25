package main

import (
	"fmt"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.
-----------------------------------------------------------------
Способ 2. Зыкрываем канал и проверяем второй параметр ok. Аналогичное решение через range,
он из под капота проверяет закрыт канал или нет, если закрыт, то выходит из цикла.
*/

func worker(stopChan chan int) {
	for {
		// value - значение из канала, ok - булево значение(true - канал открыт, false - закрыт)
		value, ok := <-stopChan
		if !ok { // Если false - канал закрыли
			fmt.Println("Канал закрылся и горутина закончила работу")
			return // С помощью return завершаем горутину
		}
		fmt.Println("Горутина в работе принято значение: ", value)
	}
}

func main() {
	stopChan := make(chan int) // Создаем канал который блокирует горутину до передачи в него значения
	go worker(stopChan)
	for i := 0; i <= 5; i++ {
		stopChan <- i
		time.Sleep(time.Second / 2)
	}
	// Закрываем канал и как следствие горутину worker.
	close(stopChan)
	// Ждем секунду что бы горутина воркер успела завершиться, до закрытия main()
	time.Sleep(time.Second)
}
