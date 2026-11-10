package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

type Rectangle struct {
	length float64
	breath float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.breath
}

type Triangle struct {
	side1 float64
	side2 float64
}

func (t Triangle) Area() float64 {
	return 1.0 / 2.0 * t.side1 * t.side2
}

func main() {

	var fig1 Rectangle = Rectangle{
		length: 10,
		breath: 25,
	}

	var fig2 Circle = Circle{
		radius: 7,
	}

	var fig3 Triangle = Triangle{
		side1: 10,
		side2: 3,
	}

	fmt.Println(fig1.Area())
	fmt.Println(fig2.Area())
	fmt.Println(fig3.Area())

}
