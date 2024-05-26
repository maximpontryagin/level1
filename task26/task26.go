package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая проверяет, что все символы в строке
уникальные (true — если уникальные, false etc). Функция проверки должна быть
регистронезависимой.
Например:
abcd — true
abCdefAaf — false
aabcd — false
-----------------------------------------------------------------
Использую мапу для хранения уже проверенных символов.
*/

// функция проверки уникальности символов в строке
func isUnique(input string) bool {
	// создаем набор для отслеживания встреченных символов
	charSet := make(map[rune]bool)

	// проходим по каждому символу строки
	for _, char := range strings.ToLower(input) {
		if _, exists := charSet[char]; exists {
			// если символ уже есть в наборе, значит он не уникален
			return false
		}
		// добавляем символ в набор
		charSet[char] = true
	}

	// если все символы уникальны
	return true
}

func main() {
	// тестовые примеры
	tests := []string{"abcd", "abCdefAaf", "aabcd"}

	for _, test := range tests {
		fmt.Printf("%s - %t\n", test, isUnique(test))
	}
}
