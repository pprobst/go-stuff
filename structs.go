package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

type writer struct {
	person
	famous bool
}

func main() {
	p1 := person{
		first: "Fyodor",
		last:  "Dostoyevsky",
		age:   35,
	}
	fmt.Println(p1.first, p1.last, "| age", p1.age)

	w1 := writer{
		person: p1,
		famous: true,
	}
	fmt.Println(w1.person.first, w1.person.last, w1.famous)

	// Anonymous struct
	p2 := struct {
		first string
		last  string
		age   int
	}{
		first: "Immanuel",
		last:  "Kant",
		age:   39,
	}
	fmt.Println(p2)
}
