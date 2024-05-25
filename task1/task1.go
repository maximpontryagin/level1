package main

import (
	"errors"
	"fmt"
)

/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры
Human (аналог наследования).
*/

// Объявление структуры человека
type Human struct {
	name       string
	profession string
	age        int
}

// Функция для создания нового объекта Human
func NewHuman(name, profession string, age int) (*Human, error) {
	// Проверка валидности переданных данных
	if name == "" {
		return nil, errors.New("пустое поле имени")
	}
	if profession == "" {
		return nil, errors.New("пустое поле профессии")
	}
	if age < 0 {
		return nil, errors.New("возраст человека не может быть отрицательным")
	}
	return &Human{name, profession, age}, nil
}

// Метод структуры Human, проверяющий является ли наш Human разработчиком
func (h *Human) AboutHuman() bool {
	return h.profession == "developer"
}

// Объявление структуры Action
type Action struct {
	// Встраивание (аналог наследования) - композиция.
	// Встраиваем указатель на объект типа Human.
	*Human
}

func main() {
	// Создаем новый объект человека
	human, err := NewHuman("Maxim", "developer", 23)
	if err != nil {
		fmt.Print(err)
		return
	}
	// Создаем переменную типа Action
	action := Action{Human: human}
	// Проверяем, является ли Human структуры Action разработчиком
	fmt.Print(action.AboutHuman()) // В нашем случем: true
}
