package main

import (
    "fmt"
)

func main() {
    even := make(chan int)
    odd := make(chan int)
    quit := make(chan bool)

    // send
    go send(even, odd, quit)

    // receive
    receive(even, odd, quit)
}

// We generate some values and put them into the respective channels
func send(e, o chan<- int, q chan<- bool) {
    for i := 0; i < 100; i++ {
        if i % 2 == 0 {
            e <- i
        } else {
            o <- i 
        } 
    }
    close(q)
}

// We take the values out of their respective channels
func receive(e, o <-chan int, q <-chan bool) {
    for {
        select {
        case v := <-e:
            fmt.Println("from the even channel:", v)
        case v := <-o:
            fmt.Println("from the odd channel:", v)
        // case _, _ = <-q:
        case <-q:
            return
        }
    }
}
