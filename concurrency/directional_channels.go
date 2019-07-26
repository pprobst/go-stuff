package main

import (
	"fmt"
)

func main() {
	r := make(<-chan int) // receiver channel
	s := make(chan<- int) // sender channel

	fmt.Printf("r: %T\n", r)
	fmt.Printf("s: %T\n", s)

	c := make(chan int)

	// send
	go foo(c)

	// receive
	bar(c)
}

// send
func foo(c chan<- int) {
	c <- 42
}

// receive
func bar(c <-chan int) {
	fmt.Println(<-c)
}
