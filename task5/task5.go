package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в
канал, а с другой стороны канала — читать. По истечению N секунд программа
должна завершаться.
-----------------------------------------------------------------
Использую context.WithTimeout для того, что бы задать время через которое он должен закрыться.
Т.е. в этом контексте задается время через которое программа должна завершиться.
*/

func main() {
	var wg sync.WaitGroup
	var N time.Duration    // количетсво секунд после которых программа завершится
	_, err := fmt.Scan(&N) // считывание с консоли времени до завершения программы
	if err != nil {
		return
	}

	// Создаем канал для передачи данных
	process := make(chan int)

	// Создаем контекст, который завершится после N секунд или функции отмены(cancel()) и отпарвит сигнал ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), N*time.Second)
	defer cancel()

	wg.Add(1)
	go func(process chan int) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done(): // Получаем сигнал из контекста о завершении программы
				fmt.Println("Завершает программу")
				return
			case values, ok := <-process: // Читаем значения из канала
				if !ok {
					return
				}
				fmt.Println(values)
			}
		}
	}(process)

	// Отправляем данные в канал process пока не завершаем программу
	wg.Add(1)
	go func() {
		i := 0
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("завершил программу по отправке значений в воркеры")
				return
			default:
				process <- i
				i++
				time.Sleep(50 * time.Millisecond) // Задержка для наглядности
			}
		}
	}()

	<-ctx.Done() // Ждем сигнал о завершении контекста
	fmt.Println("Получен сигнал завершения, завершаем работу...")
	close(process) // Закрываем канал process
	wg.Wait()      // Ждем завершения работы всех горутин
	fmt.Println("Работа завершена")
}
