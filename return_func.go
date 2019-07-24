package main

import (
	"fmt"
)

func main() {
	f1 := foo()
	fmt.Println(f1)

	f2 := func() int {
		return 1997
	}()
	fmt.Printf("%T\n", f2)

	f3 := bar()
	fmt.Printf("%T\n", f3)
	n := f3()
	fmt.Println(n)

    fmt.Println("bar()():", bar()())
}

func foo() string {
	return "Okaeri anata"
}

func bar() func() int {
	return func() int {
		return 1997
	}
}
