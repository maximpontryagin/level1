package main

import (
	"fmt"
	"time"
)

/*
Реализовать собственную функцию sleep
-----------------------------------------------------------------
Способ 1. Использование time.After
*/

func Sleep(seconds int) {
	// Создаем канал который будет блокировать Sleep на заданное количество секунд
	// по истичению времени в <- придет сигнал, что можно продолжать программу
	<-time.After(time.Second * time.Duration(seconds))
}

func main() {
	fmt.Println("Программа запустилась и уснула на 3 секунды")
	Sleep(3)
	fmt.Println("3 секунды прошли, программа проснулась!")
}