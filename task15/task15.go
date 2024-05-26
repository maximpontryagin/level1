package main

import "fmt"

/*
К каким негативным последствиям может привести данный фрагмент кода, и как
это исправить? Приведите корректный пример реализации.

var justString string
func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}
func main() {
	someFunc()
}
-----------------------------------------------------------------
Поскольку в функции someFunc создается огромная строка, а затем из неё берется подстрока,
исходная огромная строка остается в памяти, т.к. срез строки сохраняет ссылку на исходный массив байт.
В итоге, хоть нам и нужны только первые 100 символов, в памяти сохраняется вся большая строка.
createHugeString(1 << 10) создает строку длиной 1024 байт, 1 << 10 сдвигает единицу на 10 бит влево.
*/

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	// Преобразование v[:100] в []byte и обратно в string заставляет Go
	// создать новый массив байтов, который не будет ссылаться на исходную строку.
	justString = v[:100]                    // Берем первые 100 символов
	justString = string([]byte(justString)) // Копируем нужные 100 символов в новую строку
	fmt.Println(justString)
}

func createHugeString(size int) string {
	HugeRune := make([]rune, size) // Срез рун
	BaseString := []rune("T")      // Срез из руны символа T. Руна - unicode символов

	for i := range HugeRune {
		HugeRune[i] = BaseString[0] // Заполняем огромный слайс рун (для примера руной символа - "T")
	}
	HugeString := string(HugeRune) // Переводим огромный слайс рун в огромную строку
	return HugeString
}

func main() {
	someFunc()
}
