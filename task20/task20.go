package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
-----------------------------------------------------------------

*/

func main() {
	in := bufio.NewReader(os.Stdin)
	str, _ := in.ReadString('\n') // Считываем строку из консоли
	str = str[:len(str)-2]        // Удаляение 2ух последних символов /r - возврат каретки и /n - перенос строки

	list := strings.Split(str, " ") // Разбиваем исходную строку в список через пробел

	var res string // Результирующая строка
	for i := len(list) - 1; i >= 0; i-- {
		res += list[i]
		if i != 0 {
			res += " " // Если слово не последнее, то добавляем пробел между словами.
		}
	}
	fmt.Println(res)
}
