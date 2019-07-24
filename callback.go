package main

import (
	"fmt"
)

func main() {
	t := evenSum(sum, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}...)
	fmt.Println(t)
}

func sum(x ...int) int {
	n := 0
	for _, v := range x {
		n += v
	}
	return n
}

// callback = func as a parameter to another func
func evenSum(f func(x ...int) int, y ...int) int {
	var xi []int
	for _, v := range y {
		if v%2 == 0 {
			xi = append(xi, v)
		}
	}
	total := f(xi...) // the callback is used here
	return total
}
