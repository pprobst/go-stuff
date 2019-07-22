package main

import (
    "fmt"
)

// Go is a STATIC programming language!
var x int // zero value
var y = 42
var s string = "Gopher"
var raw_string string = `Oh, "my"!
                         stop

                         !"`

// We can create our own types:
type nakadashi int
var m nakadashi

func main() {
    fmt.Println("value of 'x':", x)    
    fmt.Printf("type of 'x': %T\n", x)
    fmt.Println("value of 'y':", y)    
    fmt.Printf("type of 'y': %T\n", y)
    fmt.Println("value of 's':", s)    
    fmt.Printf("type os 's': %T\n", s)
    fmt.Println("value of 'raw_string':", raw_string)    

    // short declaration operator
    z := "I'm a string!"
    fmt.Println("value of 'z':", z)    
    fmt.Printf("type of 'z': %T\n", z)

    foo()

    m = 511
    fmt.Println("value of 'm':", m)    
    fmt.Printf("type of 'm': %T\n", m)

    // Conversion (not called 'casting' in Go)
    y = int(m)
    fmt.Println("value of 'y':", y)    
    fmt.Printf("type of 'y': %T\n", y)
}

func foo() {
    fmt.Printf("'y' as binary: %b\n", y)
    fmt.Printf("'y' as hexadecimal: %#x\n", y) // %X if lower-case
    p := fmt.Sprintf("%#x\t%b\t%v", y, y, y) // %v is the default format
    fmt.Println(p)
}

