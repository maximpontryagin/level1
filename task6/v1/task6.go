package main

import (
	"fmt"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.
-----------------------------------------------------------------
Способ 1. Использование канала
*/

func worker(stopChan chan bool) {
	for {
		select {
		case <-stopChan: // Ждем пока придет сигнал из main()
			fmt.Println("Горутина остановлена")
			return
		default:
			fmt.Println("Горутина в работе...")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	stopChan := make(chan bool) // Создаем канал который блокирует горутину до передачи в него значения
	go worker(stopChan)
	// Ждем 3 секунды до закрытия горутины воркер
	time.Sleep(3 * time.Second)
	// Закрываем воркер
	stopChan <- true
}
