package main

import (
    "fmt"
)

func main() {
    // There's no "while loop" in Go!
    // Typical for:
    for i := 0; i <= 100; i++ {
        if i % 2 == 0 {
            fmt.Println("i:", i, "is pair; entering innter loop...") 
            for j := 0; j <= 5; j++ {
                fmt.Println("\tj:", j)
            }
        } else {
            fmt.Println("i:", i, "is odd") 
        }
    }

    a := 1
    // For statement with a single condition:
    for a < 10 {
        fmt.Println("a:", a)
        a++ 
        if a == 5 {
            break 
        }
    }
    fmt.Println("a:", a)

    x := 1
    for {
        x++
        if x > 100 {
            break 
        } 

        if x % 2 != 0 {
            continue // "skip"
        }

        fmt.Println("x:", x)
    }

    printAscii()

    s := "Goethe"
    switch s {
        case "Werther":
            fmt.Println("NO!")
        case "Goethe":
            fmt.Println("YES!")
        default:
            break
    }

    fmt.Println(true && true)
    fmt.Println(true && false)
    fmt.Println(true || false)
    fmt.Println(true || true)
    fmt.Println(!true)
}

func printAscii() {
    for i := 33; i <= 122; i++ {
        fmt.Printf("%v\t%#x\t%#U\n", i, i, i) 
    }
}

