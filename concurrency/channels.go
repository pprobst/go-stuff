package main

import (
    "fmt"
)

func main() {
    // Channels allow us to pass values between Goroutines
    // CHANNELS BLOCK
    // CHANNELS BLOCK 
    c := make(chan int)

    go func(){
        c <- 42 // inserts value into the channel
    }()

    fmt.Println(<-c) // takes value off
}
