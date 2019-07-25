package main

import (
	"fmt"
)

// & gives the address
// * gives the value stored at the address

func main() {
	a := 16
	fmt.Println(a)
	fmt.Println(&a) // address in memory where 'a' is stored
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a) // type is a pointer to an int

	//var b *int = &a
	b := &a
	fmt.Println(b)   // address of 'a'
	fmt.Println(*b)  // value of 'a' (16)
	fmt.Println(*&a) // value of 'a' (16)

    // because 'b' is pointing to 'a', changing the value of '*b' is
	// actually changing the value of 'a'
	*b = 32
	fmt.Println(a)
	
	x := 0
	foo(x)
	fmt.Println(x) // x is still zero
	bar(&x)
	fmt.Println(x) // x is now 50
}

func foo(y int) {
	fmt.Println("foo:")
	fmt.Println(y)
	y = 50
	fmt.Println(y)
}

func bar(y *int) { // pointer to an int = an address
	fmt.Println("bar:")
	fmt.Println(*y)
	*y = 50
	fmt.Println(*y)
}
