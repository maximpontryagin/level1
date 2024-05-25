package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

/*
Реализовать постоянную запись данных в канал (главный поток). Реализовать
набор из N воркеров, которые читают произвольные данные из канала и
выводят в stdout. Необходима возможность выбора количества воркеров при
старте.
Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать
способ завершения работы всех воркеров.
-----------------------------------------------------------------
Контекст предоставляет единый способ передачи сигнала завершения для всех горутин.
Когда вызывается cancel(), все горутины, которые наблюдают за ctx.Done(),
получают сигнал завершения и могут корректно завершить свою работу.
Данный подход делает код более гибким и позволяет легко добавлять новые горутины,
которые могут так же реагировать на сигнал завершения через контекст.
*/

func worker(ctx context.Context, proces <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done(): // Получаем сигнал из контекста о завершении программы
			fmt.Println("worker завершил программу")
			return
		case values, ok := <-proces: // Читаем значения из канала
			if !ok {
				return
			}
			fmt.Println(values)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	var amount int              // количетсво воркеров
	_, err := fmt.Scan(&amount) // считывание с консоли воркеров
	if err != nil {
		return
	}

	// Создаем буферизированного(что бы горутины не блокировались в ожидании нового значения) канала для передачи данных
	process := make(chan int, amount)

	ctx, cancel := context.WithCancel(context.Background())

	// Запускаю горутины воркеров в цикле
	for i := 0; i < amount; i++ {
		wg.Add(1)
		go worker(ctx, process, &wg)
	}

	// Отправляем данные в канал process(в воркеры) пока не завершаем программу
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

	// Завершение по нажатию Ctrl+C
	quit := make(chan os.Signal, 1)                      // Создание канала с буффером 1
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT) // Подписка на сигналы завершения программы
	// После нажатия Ctrl+C наш main запускает процесс остановки горутин
	<-quit
	fmt.Println("Получен сигнал завершения, завершаем работу...")
	cancel()       // После Ctrl+C закрываем контекст, что бы сообщить го рутинам, что пора завершать работу
	close(process) // Закрываем канал process
	wg.Wait()      // Ждем завершения работы всех горутин
	fmt.Println("Работа завершена")
}

/*
Этот подход обеспечивает безопасное завершение программы и позволяет воркерам
корректно завершить свою работу после получения сигнала завершения. Ниже приведу пример
завершения программы без контекста, но такой подход будет менее гибким
*/

// package main

// import (
// 	"fmt"
// 	"os"
// 	"os/signal"
// 	"sync"
// 	"syscall"
// 	"time"
// )

// func worker(quit <-chan struct{}, proces <-chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for {
// 		select {
// 		case <-quit: // Чтение из закрытого канала завершения
// 			fmt.Println("worker завершил программу")
// 			return
// 		case values, ok := <-proces: // Чтение из канала данных
// 			if !ok {
// 				return
// 			}
// 			fmt.Println(values)
// 		}
// 	}
// }

// func main() {
// 	var wg sync.WaitGroup
// 	var amount int
// 	fmt.Scan(&amount)

// 	process := make(chan int, amount)
// 	quit := make(chan struct{})

// 	for i := 0; i < amount; i++ {
// 		wg.Add(1)
// 		go worker(quit, process, &wg)
// 	}

// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		i := 0
// 		for {
// 			select {
// 			case <-quit:
// 				fmt.Println("завершил программу по отправке значений в воркеры")
// 				return
// 			default:
// 				process <- i
// 				i++
// 				time.Sleep(50 * time.Millisecond)
// 			}
// 		}
// 	}()

// 	// Ожидание сигнала завершения (Ctrl+C)
// 	sigChan := make(chan os.Signal, 1)
// 	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)
// 	<-sigChan

// 	fmt.Println("Получен сигнал завершения, завершаем работу...")
// 	close(quit) // Закрытие канала для сигнализации завершения
// 	close(process)
// 	wg.Wait() // Ожидание завершения всех горутин
// 	fmt.Println("Работа завершена")
// }
