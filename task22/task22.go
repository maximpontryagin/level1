package main

import (
	"fmt"
	"math/big"
)

/*
Разработать программу, которая перемножает, делит, складывает, вычитает две
числовых переменных a,b, значение которых > 2^20.
-----------------------------------------------------------------
Применена библиотека "math/big"
*/

func main() {
	// Создаем большие числа a и b
	a := new(big.Int)
	b := new(big.Int)

	// Первый аргумент - строка представляющая число, вторая - система счисления, в данном случае десятичная
	a.SetString("1048577", 10) // 2^20 + 1
	b.SetString("2097153", 10) // 2^21 + 1

	// Переменные для хранения результатов
	sum := new(big.Int)
	diff := new(big.Int)
	prod := new(big.Int)
	quot := new(big.Int)

	// Выполняем операции
	sum.Add(a, b)  // Сложение
	diff.Sub(a, b) // Вычитание
	prod.Mul(a, b) // Умножение
	// Возникает паника при делении на ноль
	quot.Quo(a, b) // Деление

	fmt.Printf("a: %s\n", a.String())
	fmt.Printf("b: %s\n", b.String())
	fmt.Printf("Сумма: %s\n", sum.String())
	fmt.Printf("Вычитание: %s\n", diff.String())
	fmt.Printf("Умножение: %s\n", prod.String())
	fmt.Printf("Деление: %s\n", quot.String())
}
