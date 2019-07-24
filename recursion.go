package main

import (
	"fmt"
)

func main() {
	fmt.Println(factorial(5))
	fmt.Println(fact(5))
}

// No big deal, same as every language out there
// Recursive factorial
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func fact(n int) int {
	sum := 1
	for ; n > 0; n-- {
		sum *= n
	}
	return sum
}
