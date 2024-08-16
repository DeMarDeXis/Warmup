package main

import (
	"fmt"
)

// Определение интерфейса Shape
type Shape interface {
	Area() float64
}

// Реализация интерфейса для круга
type Circle struct {
	Radius float64
}

// Метод для вычисления площади круга (реализация интерфейса)
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Реализация интерфейса для прямоугольника
type Rectangle struct {
	Width  float64
	Height float64
}

// Метод для вычисления площади прямоугольника (реализация интерфейса)
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	// Создание экземпляров фигур
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 4, Height: 6}

	// Использование интерфейса для расчета площадей разных фигур
	shapes := []Shape{circle, rectangle}
	for _, shape := range shapes {
		fmt.Printf("Площадь фигуры: %.2f\n", shape.Area())
	}
}
