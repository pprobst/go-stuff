package main

import (
	"fmt"
)

func main() {
	func() {
		fmt.Println("Anonymous function")
	}()

	func(s string) {
		fmt.Println("Day of my birthday:", s)
	}("06/16")
}
