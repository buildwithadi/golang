package main

import "fmt"

type Rectangle struct {
	length float64
	width  float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.width)
}

func (r Rectangle) isSquare() bool {
	if r.length == r.width {
		return true
	}
	return false
}

func (r *Rectangle) Switch() bool {
	temp := r.length
	r.length = r.width
	r.width = temp

	return true
}

func main() {

	fig := Rectangle{
		length: 10,
		width:  20,
	}

	fmt.Println(fig.isSquare())
	fmt.Println(fig.Area())
	fmt.Println(fig.Perimeter())
	fmt.Println(fig.Switch())
	fmt.Println(fig.length)
}
