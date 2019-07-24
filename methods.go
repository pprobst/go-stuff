package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

type writer struct {
	person
	famous bool
}

// func (r receiver) identifier(parameters) (return(s)) { body }
// A method in Go:
func (w writer) speak() {
	fmt.Print("I am ", w.first, " ", w.last)
	if w.famous {
		fmt.Print(", and I'm famous!\n")
	} else {
		fmt.Print("... and I'm not... famous...\n")
	}
}

func main() {
	w1 := writer{
		person: person{
			"Franz",
			"Kafka",
		},
		famous: true,
	}

	w2 := writer{
		person: person{
			"Mimi",
			"Nini",
		},
		famous: false,
	}

	fmt.Println(w1)
	w1.speak()
	fmt.Println(w2)
	w2.speak()
}
