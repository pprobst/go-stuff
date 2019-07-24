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

func (w writer) speak() {
	fmt.Print("I am ", w.first, " ", w.last)
	if w.famous {
		fmt.Print(", and I'm famous!\n")
	} else {
		fmt.Print("... and I'm not... famous...\n")
	}
}

func (p person) speak() {
	fmt.Print("I am ", p.first, " ", p.last)
}

// Interface in Go:
type human interface {
	// Any type that has "speak()" is also of type 'human'
	speak()
}

func bar(h human) {
	switch h.(type) { // switch on type
	case person:
        fmt.Println("person passed into bar:", h.(person).first) // assertion
	case writer:
        fmt.Println("writer passed into bar:", h.(writer).first)
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

	p1 := person{
		first: "Nani",
		last:  "Mo",
	}

	fmt.Println(p1)

	bar(w1)
	bar(w2)
	bar(p1)
}
