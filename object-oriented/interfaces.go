package main

import (
	"fmt"
	"math"
)

type Square1 struct {
	Length float64
}

func (s Square1) Area() float64 {
	return s.Length * s.Length
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func sumAreas(shapes []Shape) float64 {
	total := 0.0

	for _, shape := range shapes {
		total += shape.Area()
	}

	return total
}

// Defines set of methods that a type should implement
type Shape interface {
	Area() float64
}

func main() {
	s := Square1{20}
	fmt.Println(s.Area())

	c := Circle{10}
	fmt.Println(c.Area())

	shapes := []Shape{s, c}
	sa := sumAreas(shapes)
	fmt.Println(sa)
}
