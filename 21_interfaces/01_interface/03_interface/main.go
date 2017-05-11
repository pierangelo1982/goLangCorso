package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

// another shape
type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (z square) area() float64 {
	return z.side * z.side
}

// which implements the shape interface
// che implementa l'interfaccia shape
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(z shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	s := square{10}
	c := circle{5}
	info(s)
	info(c)
}
