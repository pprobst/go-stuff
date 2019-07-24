package main

import (
	"fmt"
	m "math"
)

type circle struct {
	radius float64
}

type square struct {
	length float64
}

func (c circle) area() float64 {
	return m.Pi * (c.radius * c.radius)
}

func (s square) area() float64 {
	return s.length * s.length
}

type shape interface {
	area() float64
}

func info(sp shape) {
	fmt.Println(sp.area())
}

func main() {
	c1 := circle{
		radius: 7.56,
	}
	sq1 := square{
		length: 20.33,
	}

	info(c1)
	info(sq1)
}
