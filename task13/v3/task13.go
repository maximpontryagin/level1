package main

import "fmt"

/*
Поменять местами два числа без создания временной переменной.
-----------------------------------------------------------------
Использование множественного присваивания.
*/

func main() {
	a := 5
	b := 10

	fmt.Printf("Было a = %d, b = %d\n", a, b)

	a, b = b, a

	fmt.Printf("Стало a = %d, b = %d\n", a, b)
}