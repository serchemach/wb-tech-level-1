package main

import (
	"fmt"
	"math"
)

// В го нет как таковой инкапсуляции на уровне одного пакета,
// поэтому в одном файле её показать не особо возможно.
// Но если использовать эту структуру в другом пакете, то х и у будут
// инкапсулированы, поскольку начинаются с маленькой буквы
type Point struct {
	x, y float64
}

// Как таковых, конструкторов тоже нет, но обычно для таких целей делают
// функции с названием или преффиксом New
func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

func (p1 *Point) Distance(p2 *Point) float64 {
	return math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2))
}

func main() {
	p1 := NewPoint(0, 0)
	p2 := NewPoint(1, 1)

	fmt.Printf("The distance between points %#v and %#v is %f\n", p1, p2, p1.Distance(p2))
}
