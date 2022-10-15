package main

import (
	"fmt"
)

func main() {
	x := 10

	if x > 5 {
		fmt.Println("x is big")
	}

	if x > 100 {
		fmt.Println("x is very big")
	} else {
		fmt.Println("x is not that big")
	}

	if x > 5 && x < 15 {
		fmt.Println("x is just right")
	}

	if x < 20 || x > 30 {
		fmt.Println("x is out of range")
	}

	a, b := 11.0, 20.0

	if frac := a / b; frac > 0.5 {
		fmt.Println("a is more than half of b")
	}

	n := 2

	switch n {
	case 1:
		fmt.Printf("one")
	case 2:
		fmt.Printf("two\n")
	case 3:
		fmt.Printf("three")
	default:
		fmt.Printf("many")
	}

	// naked switch
	switch {
	case n > 100:
		fmt.Printf("n is very big")
	case n > 10:
		fmt.Printf("n is big")
	default:
		fmt.Println("n is small")
	}
}
