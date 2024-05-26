package main

import (
	"fmt"
	"math"
)

/*
Разработать программу нахождения расстояния между двумя точками, которые
представлены в виде структуры Point с инкапсулированными параметрами x,y и
конструктором.
-----------------------------------------------------------------

*/

type Point struct {
	x float64
	y float64
}

// Функция для создания новой точки
func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

// Функция поиска расстояния между точками
func DistanceBetweenPoints(point1, point2 *Point) float64 {
	dx := point1.x - point2.x
	dy := point1.y - point2.y
	distance := math.Sqrt(dx*dx + dy*dy)
	return distance
}

func main() {
	point1 := NewPoint(1, 1)
	point2 := NewPoint(2, 2)
	distance := DistanceBetweenPoints(point1, point2)
	fmt.Println("Расстояние между точками 1 и 2: ", distance)
}
