package main

import (
    "fmt"
)

func main() {
    c := make(chan int, 1)

    c <- 42
    // c <- 43 this will block!
    // Channels have to pass/receive the value at the same time!

    fmt.Println(<-c)

    c2 := make(chan int, 2)
    c2 <- 16
    c2 <- 32

    fmt.Println(<-c2) // pulls off 16
    fmt.Println(<-c2) // pulls off 32
}
