package main

import (
	"context"
	"fmt"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.
-----------------------------------------------------------------
Способ 2. Создаем контекст, закрываем его, в горутине worker получаем
сигнал о закрытии контекста -> вызываем return для закрытия горутины.
*/

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): // Ловим сигнал, что закрылся контекст
			fmt.Println("Горутина остановлена")
			return // Выходим из горутины после закрытия контекста в main()
		default:
			fmt.Println("Горутина в работе...")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	// Создаем контекст с функцией cancel для передачи сигнала о его закрытии.
	ctx, cancel := context.WithCancel(context.Background())
	go worker(ctx) // запускаем горутина worker

	time.Sleep(3 * time.Second)
	// Закрываем контекст ctx => по ctx.Done() пойдет сигнал,
	// что контекст зарылся(это будет сигналом на закрытие горутины)
	cancel()
	// Ждем секунду, что бы горутина woeker успела завершиться, до закрытия main()
	time.Sleep(1 * time.Second)
}
