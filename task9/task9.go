package main

import (
	"fmt"
	"sync"
)

/*
Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из
массива, во второй — результат операции x*2, после чего данные из второго
канала должны выводиться в stdout.
-----------------------------------------------------------------
Данные через каналы будут передаваться через 2 горутины, после чего поступают
в третью горутину и выводятся в stdout.
*/

func main() {
	var wg sync.WaitGroup
	array := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	inCh := make(chan int)
	OutCh := make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, val := range array {
			inCh <- val
		}
		close(inCh)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range inCh {
			OutCh <- val * 2
		}
		close(OutCh)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range OutCh {
			fmt.Println(val)
		}
	}()

	wg.Wait()
}
