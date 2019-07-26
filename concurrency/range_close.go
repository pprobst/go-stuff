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
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	// receive
	for v := range c {
		fmt.Println(v)
	}
}
