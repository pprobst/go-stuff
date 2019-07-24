package main

import (
    "fmt"
)

func main() {
    f := func(x int) {
        fmt.Println("Func expression,", x) 
    }
    f(0)
}
