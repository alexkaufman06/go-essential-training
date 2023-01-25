package main

import (
	"fmt"
	"log"
)

// Define a square structure. It should have three fields, X and a Y, which are integers, and a length, which is also an
// integer. Create a new function to create a new square and add two methods. One, Move, which will move the square by
// Delta X and Delta Y, and an area that will return the area of the square. In the main, you should be able to create a
// new square with new square. Move it to a new location. You can print it out and also print out the area.

type Square struct {
	X      int
	Y      int
	Length int
}

func NewSquare(x int, y int, l int) (*Square, error) {
	if l <= 0 {
		return nil, fmt.Errorf("Length value must be greater than 0")
	}

	s := Square{
		X:      x,
		Y:      y,
		Length: l,
	}
	return &s, nil
}

func (s *Square) Move(dx int, dy int) {
	s.X += dx
	s.Y += dy
}

func (s Square) Area() int {
	return s.Length * s.Length
}

func main() {
	s, err := NewSquare(1, 1, 4)
	if err != nil {
		log.Fatalf("ERROR: can't create square")
	}

	fmt.Printf("%+v\n", s)
	s.Move(1, 1)
	fmt.Println("I moved the square to:")
	fmt.Printf("%+v\n", s)
	fmt.Println("The area of the square is:")
	fmt.Println(s.Area())
}
