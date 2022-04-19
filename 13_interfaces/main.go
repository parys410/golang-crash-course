package main

import (
	"fmt"
	"math"
)

// Define interface
type Shape interface {
	area() float64
}

type Circle struct {
	X,Y,Radius float64
}

type Rectangle struct {
	Width,Height float64
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func (r Rectangle) area() float64 {
	return r.Width * r.Height
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle {
		X: 0,
		Y: 0,
		Radius: 5,
	}

	rectangle := Rectangle {
		Width: 10,
		Height: 5,
	}

	fmt.Printf("Circle Area: %f\n", getArea(circle))
	fmt.Printf("Rectangle Area: %f\n", getArea(rectangle))
}

