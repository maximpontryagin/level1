package main

import "fmt"

/*
Разработать программу, которая в рантайме способна определить тип
переменной: int, string, bool, channel из переменной типа interface{}.
-----------------------------------------------------------------
Переключатель типов позволяет выбирать между типами.
v := value.(type): - выполняет утверждение типа (type assertion) и одновременно
определяет тип значения, содержащегося в переменной value типа interface{}.
ivalue.(type): Специальный синтаксис, используемый в конструкциях type switch,
который позволяет проверить конкретный тип значения, находящегося в интерфейсе value.
*/

func identifyType(value interface{}) {
	switch v := value.(type) {
	case int:
		fmt.Printf("тип = \"%T\", значение = %v", v, v)
		fmt.Println()
	case string:
		fmt.Printf("тип = \"%T\", значение = %v", v, v)
		fmt.Println()
	case bool:
		fmt.Printf("тип = \"%T\", значение = %v", v, v)
		fmt.Println()
	case chan interface{}:
		fmt.Printf("тип = \"%T\", значение = %v", v, v)
		fmt.Println()
	default:
		fmt.Println("unknown type")
	}
}

func main() {
	// Примеры переменных разных типов
	var a int = 42
	var b string = "hello"
	var c bool = true
	var d chan interface{} = make(chan interface{})

	// Определение типа каждой переменной
	identifyType(a)
	identifyType(b)
	identifyType(c)
	identifyType(d)

	// Пример с переменной неизвестного типа
	var e float64 = 3.14
	identifyType(e)
}
