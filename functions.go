package main

import (
    "fmt"
)

func main() {
    foo()
    bar("Bar")
    s1 := woo("Woo")
    fmt.Println(s1)
    f, l := zoo ("Alan", "Turing")
    fmt.Println(f)
    fmt.Println(l)
    sum := variadicParameter(1,2,3,4,5,6,7,8,9)
    fmt.Println(sum)
}

func foo() {
    fmt.Println("Foo")
}

func bar (s string) {
    fmt.Println(s)
}

func woo (s string) string {
    return fmt.Sprint(s)
}

func zoo (f, l string) (string, bool) {
    a := fmt.Sprint(f, " ", l, ` says "Hello."`) 
    b := true
    return a, b
}

func variadicParameter(x ...int) (sum int) {
    fmt.Println(x)
    fmt.Printf("%T\n", x)

    sum = 0
    for _, v := range x {
        sum += v 
    }

    return
}
