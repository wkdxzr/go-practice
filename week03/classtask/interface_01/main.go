package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Triangle struct {
	SideA, SideB, SideC float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.SideA * t.SideB
}
func (t Triangle) Perimeter() float64 {
	return t.SideA + t.SideB + t.SideC
}

func main() {
	shapes := []Shape{
		Circle{Radius: 5},
		Rectangle{Width: 4, Height: 6},
		Triangle{SideA: 3, SideB: 4, SideC: 5},
	}

	for _, shape := range shapes {
		fmt.Printf("图形面积: %.2f, 图形周长: %.2f\n", shape.Area(), shape.Perimeter())
	}
}
