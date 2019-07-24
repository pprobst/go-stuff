package main

import (
    "fmt"
)

func main() {
    xi := []int{1,2,3,4,5}
    x := sum(xi...) // "unfurling" a slice with '...'
    fmt.Println(x)
}

func sum(x ...int) (sum int) {
    for _, v := range(x) {
        sum += v   
    }

    return
}
