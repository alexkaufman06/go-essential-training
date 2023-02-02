package main

import "fmt"

func main() {
	ch := make(chan int)
	//ch <- 353 // send (fatal error: all goroutines are asleep - deadlock!)

	go func() {
		// Send number of the channel
		ch <- 353
	}()

	val := <-ch // receieve
	fmt.Printf("got %d\n", val)

	fmt.Println("---------")
	// https://www.linkedin.com/learning/go-essential-training-16567666/channels?autoSkip=true&autoplay=true&resume=false&u=30495498
}
