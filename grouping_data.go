package main

import (
    "fmt"
)

func main() {
    // Array
    // -----
    var x [5]int
    fmt.Println(x)
    x[3], x[4] = 16, 32
    fmt.Println(x)
    fmt.Println(len(x))

    // Slice (made on top of arrays; dynamic)
    // -----
    // composite literal --
    // x := type{values}
    y := []int{1, 2, 3, 4, 5}
    fmt.Println(y)

    // for range
    for i, e := range y {
       fmt.Println(i, e) 
    }
    
    // slicing a slice
    fmt.Println(y[3:])

    // appending to a slice
    z := []int{100, 101, 102}
    y = append(y, 10, 11, 12, 13)
    y = append(y, z...)
    fmt.Println(y)

    // deleting from a slice
    y = append(y[:2], y[4:]...)
    fmt.Println(y)

    // make
    // used for optimization if you already know the size of your slice
    m := make([]int, 10, 100)
    m[0] = 2
    m[9] = 100
    fmt.Println(m)
    fmt.Println(len(m))
    fmt.Println(cap(m))
    m = append(m, 101)
    fmt.Println(m)
    fmt.Println(len(m)) // goes up after the append!
    fmt.Println(cap(m))

    // multi-dimensional slice
    a := []string{"A", "AA", "Lemon"}
    b := []string{"B", "BB", "Chocolate"}
    ab := [][]string{a, b}
    fmt.Println(ab)

    // Map
    // ---
    t := map[string]int {
            "Mom": 46,
            "Me":  22,
    }
    fmt.Println(t)
    fmt.Println(t["Me"])
    if val, ok := t["Dad"]; ok {
        fmt.Println(val, ok) // won't print because Dad is not mapped
    }

    t["Grandfather"] = 87 // new entry
    for k, v := range t {
        fmt.Println(k, v) 
    }

    // deleting an entry from a map
    delete(t, "Grandfather")
    fmt.Println("after deleting grandfather:", t)
}
