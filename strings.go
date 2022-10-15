package main

import (
	"fmt"
)

func main() {
	book := "The Grapes of Wrath"
	fmt.Println(book)
	fmt.Println(len(book))
	fmt.Println("book[0] = %v (type %T)", book[0], book[0])
	// uint8 is a byte

	// strings are immutable in Go
	//book[0] = 1

	// Slice (start, end), 0 based, 1/2 empty range
	fmt.Println(book[4:11])

	// Slice (no end)
	fmt.Println(book[4:])

	// Slice (no start)
	fmt.Println(book[:4])

	// Use + to concatenate strings
	fmt.Println("t" + book[1:])

	// Unicode
	fmt.Println("It was ½ off!")

	// Multi line
	poem := `
		The road goes ever on
		Down from the door
		To the ocean floor
		...
    `
	fmt.Println(poem)
}
