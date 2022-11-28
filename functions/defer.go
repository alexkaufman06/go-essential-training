package main

import (
	"fmt"
)

func main() {
	worker()
}

func worker() {
	r1, err := acquire("A")
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	defer release(r1) // defer only calls on function exit

	r2, err := acquire("B")
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	defer release(r2) // defers are called in reverse order (First B then A)

	fmt.Println("Worker")
}

func acquire(name string) (string, error) {
	return name, nil
}

func release(name string) {
	fmt.Printf("Cleaning up %s\n", name)
}