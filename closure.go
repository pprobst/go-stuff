package main

import (
    "fmt"
)

func main() {
    a := incrementor()
    b := incrementor()
    // a's x != b's x
    fmt.Println("a:", a())
    fmt.Println("a:", a())
    fmt.Println("a:", a())
    fmt.Println("a:", a())
    fmt.Println("b:", b())
    fmt.Println("b:", b())
}

func incrementor() func() int {
    var x int
    return func() int { // "enclosing" x
        x++
        return x
    }
}
