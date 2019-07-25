package main

import (
    "fmt"
    "sort"
)

func main() {
    xi := []int{1, 0, -2, 8, 100, 99, 54, 500, 499}
    xs := []string{"Peter", "Sadamoto", "Kurosawa", "Ann"}

    fmt.Println(xi)
    sort.Ints(xi)
    fmt.Println(xi)

    fmt.Println(xs)
    sort.Strings(xs)
    fmt.Println(xs)
}
