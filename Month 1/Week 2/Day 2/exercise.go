package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2*r.width + 2*r.height
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func PrintShapeInfo(s Shape) {
	// if _, ok := s.(Circle); ok {
	// 	fmt.Printf("The Area of the Circle is: %f\n The Perimeter of the Circle is: %f\n", s.Area(), s.Perimeter())
	// }
	// if _, ok := s.(Rectangle); ok {
	// 	fmt.Printf("The Area of the Rectangle is: %.1f\n The Perimeter of the Rectangle is: %.1f\n", s.Area(), s.Perimeter())
	// }
	fmt.Printf("The Area of the shape is: %f\n The Perimeter of the shape is: %f\n", s.Area(), s.Perimeter())
}

func main() {
	r := Rectangle{width: 2, height: 5}
	c := Circle{radius: 5}
	PrintShapeInfo(r)
	PrintShapeInfo(c)
}
